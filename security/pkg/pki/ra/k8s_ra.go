// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ra

import (
	"bytes"
	"fmt"
	"strings"
	"sync"
	"time"

	cert "k8s.io/api/certificates/v1"
	clientset "k8s.io/client-go/kubernetes"

	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/log"
	"istio.io/istio/security/pkg/k8s/chiron"
	"istio.io/istio/security/pkg/pki/ca"
	raerror "istio.io/istio/security/pkg/pki/error"
	"istio.io/istio/security/pkg/pki/util"
)

// KubernetesRA integrated with an external CA using Kubernetes CSR API
type KubernetesRA struct {
	csrInterface clientset.Interface

	// Set if the ./etc/external-ca-cert/root-cert.pem is mounted
	keyCertBundle *util.KeyCertBundle

	raOpts *IstioRAOptions

	// Key is the signer name
	// Value is Root CAs (PEM list)
	caCertificatesFromMeshConfig map[string]string

	// certSignerDomain is based on CERT_SIGNER_DOMAIN env variable
	// it is concatenanted with CertSigner metadata from the request to get the key for
	// the root certificates in caCertificatesFromMeshConfig
	certSignerDomain string

	// mutex protects the R/W to caCertificatesFromMeshConfig.
	mutex sync.RWMutex
}

var pkiRaLog = log.RegisterScope("pkira", "Istiod RA log")

// NewKubernetesRA : Create a RA that interfaces with K8S CSR CA
func NewKubernetesRA(raOpts *IstioRAOptions) (*KubernetesRA, error) {
	istioRA := &KubernetesRA{
		csrInterface: raOpts.K8sClient,
		raOpts:       raOpts,
		// CertSignerDomain is based on CERT_SIGNER_DOMAIN env variable
		certSignerDomain:             raOpts.CertSignerDomain,
		caCertificatesFromMeshConfig: make(map[string]string),
	}
	return istioRA, nil
}

// kubernetesSign will use the Kubernetes CSR API to sign - the call may include a 'certSigner' from the incoming
// gRPC metadata 'CertSigner', allowing the use of different signers for different workloads.
// This feature requires 'certSignerDomain' to be set -  based on CERT_SIGNER_DOMAIN env variable - which is used as
// prefix. If CERT_SIGNER_DOMAIN is empty the only cert signer is set via K8S_SIGNER, which is also the default if
// 'CertSigner' metadata is not set.
//
// For example, K8S_SIGNER=issuers.cert-manager.io/sandbox.my-issuer can be used as default, and
// CERT_SIGNER_DOMAIN=issuers.cert-manager.io will allow users to pick different issuers.
// Istio does not check the value of the issuer.
//
// Istiod will do the validation and auth - and auto-approve the certificate - so in the previous example a user
// will be able to use any issuer from any namespace. Use with caution.
func (r *KubernetesRA) kubernetesSign(csrPEM []byte, caCertFile string, certSigner string,
	requestedLifetime time.Duration,
) ([]byte, error) {
	certSignerDomain := r.certSignerDomain
	if certSignerDomain == "" && certSigner != "" {
		return nil, raerror.NewError(raerror.CertGenError, fmt.Errorf("certSignerDomain is required for signer %s", certSigner))
	}
	if certSignerDomain != "" && certSigner != "" {
		certSigner = certSignerDomain + "/" + certSigner
	} else {
		certSigner = r.raOpts.CaSigner
	}
	usages := []cert.KeyUsage{
		cert.UsageDigitalSignature,
		cert.UsageKeyEncipherment,
		cert.UsageServerAuth,
		cert.UsageClientAuth,
	}
	certChain, _, err := chiron.SignCSRK8s(r.csrInterface, csrPEM, certSigner, usages, "", caCertFile, true, false, requestedLifetime)
	if err != nil {
		return nil, raerror.NewError(raerror.CertGenError, err)
	}
	return certChain, err
}

// Sign takes a PEM-encoded CSR and cert opts, and returns a certificate signed by k8s CA.
// It returns the leaf certificate only.
func (r *KubernetesRA) Sign(csrPEM []byte, certOpts ca.CertOpts) ([]byte, error) {
	_, err := preSign(r.raOpts, csrPEM, certOpts.SubjectIDs, certOpts.TTL, certOpts.ForCA)
	if err != nil {
		return nil, err
	}
	certSigner := certOpts.CertSigner

	return r.kubernetesSign(csrPEM, r.raOpts.CaCertFile, certSigner, certOpts.TTL)
}

// SignWithCertChain is similar to Sign but returns the leaf cert and the entire cert chain.
// root cert comes from two sources, order matters:
// 1. Specified in mesh config
// 2. Extract from the cert-chain signed by the CSR signer.
// If no root cert can be found from either of the two sources, error returned.
// There are several possible situations:
// 1. root cert is specified in mesh config and is empty in signed cert chain, in this case
// we verify the signed cert chain against the root cert from mesh config and append the
// root cert into the cert chain.
// 2. root cert is specified in mesh config and also can be extracted in signed cert chain, in this
// case we verify the signed cert chain against the root cert from mesh config and append it
// into the cert chain if the two root certs are different. This is typical when
// the returned cert chain only contains the intermediate CA.
// 3. root cert is not specified in mesh config but can be extracted in signed cert chain, in this case
// we verify the signed cert chain against the root cert and return the cert chain directly.
func (r *KubernetesRA) SignWithCertChain(csrPEM []byte, certOpts ca.CertOpts) ([]string, error) {
	cert, err := r.Sign(csrPEM, certOpts)
	if err != nil {
		return nil, err
	}
	chainPem := r.GetCAKeyCertBundle().GetCertChainPem()
	if len(chainPem) > 0 {
		cert = append(cert, chainPem...)
	}
	respCertChain := []string{string(cert)}
	var possibleRootCert, rootCertFromMeshConfig, rootCertFromCertChain []byte
	certSigner := r.certSignerDomain + "/" + certOpts.CertSigner

	if len(r.GetCAKeyCertBundle().GetRootCertPem()) == 0 {
		// If the key bundle does not have a root - missing config - we use the last
		// element in the returned chain as root.

		rootCertFromCertChain, err = util.FindRootCertFromCertificateChainBytes(cert)
		if err != nil {
			pkiRaLog.Infof("failed to find root cert from signed cert-chain (%v)", err.Error())
		}
		rootCertFromMeshConfig, err = r.GetRootCertFromMeshConfig(certSigner)
		if err != nil {
			pkiRaLog.Infof("failed to find root cert from mesh config (%v)", err.Error())
		}
		if rootCertFromMeshConfig != nil {
			possibleRootCert = rootCertFromMeshConfig
		} else if rootCertFromCertChain != nil {
			possibleRootCert = rootCertFromCertChain
		}
		if possibleRootCert == nil {
			return nil, raerror.NewError(raerror.CSRError, fmt.Errorf("failed to find root cert from either signed cert-chain or mesh config"))
		}
		if verifyErr := util.VerifyCertificate(nil, cert, possibleRootCert, nil); verifyErr != nil {
			return nil, raerror.NewError(raerror.CSRError, fmt.Errorf("root cert from signed cert-chain is invalid (%v)", verifyErr))
		}
		if !bytes.Equal(possibleRootCert, rootCertFromCertChain) {
			respCertChain = append(respCertChain, string(possibleRootCert))
		}
	}
	return respCertChain, nil
}

// GetCAKeyCertBundle returns the KeyCertBundle for the CA, if ./etc/external-ca-cert/root-cert.pem is mounted
func (r *KubernetesRA) GetCAKeyCertBundle() *util.KeyCertBundle {
	return r.keyCertBundle
}

func (r *KubernetesRA) SetCACertificatesFromFile(roots string) error {
	keyCertBundle, err := util.NewKeyCertBundleWithRootCertFromFile(roots)
	if err != nil {
		return raerror.NewError(raerror.CAInitFail, fmt.Errorf("error processing Certificate Bundle for Kubernetes RA"))
	}
	r.keyCertBundle = keyCertBundle

	return nil
}

func (r *KubernetesRA) SetCACertificatesFromMeshConfig(caCertificates []*meshconfig.MeshConfig_CertificateData) {
	r.mutex.Lock()
	for _, pemCert := range caCertificates {
		// TODO:  take care of spiffe bundle format as well
		cert := pemCert.GetPem()
		certSigners := pemCert.CertSigners
		if len(certSigners) != 0 {
			certSigner := strings.Join(certSigners, ",")
			if cert != "" {
				r.caCertificatesFromMeshConfig[certSigner] = cert
			}
		}
	}
	r.mutex.Unlock()
}

func (r *KubernetesRA) GetRootCertFromMeshConfig(signerName string) ([]byte, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()
	caCertificates := r.caCertificatesFromMeshConfig
	if len(caCertificates) == 0 {
		return nil, fmt.Errorf("no caCertificates defined in mesh config")
	}
	for signers, caCertificate := range caCertificates {
		signerList := strings.Split(signers, ",")
		if len(signerList) == 0 {
			continue
		}
		for _, signer := range signerList {
			if signer == signerName {
				return []byte(caCertificate), nil
			}
		}
	}
	return nil, fmt.Errorf("failed to find root cert for signer: %v in mesh config", signerName)
}
