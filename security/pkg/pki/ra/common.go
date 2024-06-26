// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ra

import (
	"fmt"
	"time"

	clientset "k8s.io/client-go/kubernetes"

	meshconfig "istio.io/api/mesh/v1alpha1"
	"istio.io/istio/pkg/slices"
	raerror "istio.io/istio/security/pkg/pki/error"
	"istio.io/istio/security/pkg/pki/util"
	caserver "istio.io/istio/security/pkg/server/ca"
)

// RegistrationAuthority : Registration Authority interface.
type RegistrationAuthority interface {
	caserver.CertificateAuthority
	// SetCACertificatesFromMeshConfig sets the CACertificates using the ones from mesh config
	SetCACertificatesFromMeshConfig([]*meshconfig.MeshConfig_CertificateData)
	// GetRootCertFromMeshConfig returns the root cert for the specific signer in mesh config
	GetRootCertFromMeshConfig(signerName string) ([]byte, error)
}

// CaExternalType : Type of External CA integration
type CaExternalType string

// IstioRAOptions : Configuration Options for the IstioRA
type IstioRAOptions struct {
	// ExternalCAType: Integration API type with external CA
	// Only ISTIOD_RA_KUBERNETES_API is supported
	// deprecated - when we add a different kind we can add the right setting.
	ExternalCAType CaExternalType

	// DefaultCertTTL: Default Certificate TTL
	DefaultCertTTL time.Duration

	// MaxCertTTL: Maximum Certificate TTL that can be requested
	MaxCertTTL time.Duration

	// CaCertFile : File containing PEM encoded CA root certificate of external CA
	// Mounted from ./etc/external-ca-cert/root-cert.pem (external-ca-cert volume)
	// or /var/run/secrets/kubernetes.io/serviceaccount/ca.crt for legacy signer
	CaCertFile string

	// CaSigner : To indicate custom CA Signer name when using external K8s CA
	CaSigner string

	// VerifyAppendCA : Whether to use caCertFile containing CA root cert to verify and append to signed cert-chain
	VerifyAppendCA bool
	// K8sClient : K8s API client
	K8sClient clientset.Interface
	// TrustDomain
	TrustDomain string
	// CertSignerDomain is based on CERT_SIGNER_DOMAIN env variable
	CertSignerDomain string
}

const (
	// ExtCAK8s : Integrate with external CA using k8s CSR API
	ExtCAK8s CaExternalType = "ISTIOD_RA_KUBERNETES_API"

	// DefaultExtCACertDir : Location of external CA certificate
	DefaultExtCACertDir string = "./etc/external-ca-cert"
)

// ValidateCSR : Validate all SAN extensions in csrPEM match authenticated identities
func ValidateCSR(csrPEM []byte, subjectIDs []string) bool {
	csr, err := util.ParsePemEncodedCSR(csrPEM)
	if err != nil {
		return false
	}
	if err := csr.CheckSignature(); err != nil {
		return false
	}

	// csrIDs is the list with all the SANs in the CSR. It ignores the types.
	csrIDs, err := util.ExtractIDs(csr.Extensions)
	if err != nil {
		return false
	}

	// For each SAN in the CSR, verify that it is part of the list of allowed certificates in the
	// subjectIDs list, which is the list of identities Istiod validated.
	// In the case of Ztunnel, this is the list from the metadata, checked by the NodeAuthorizer.
	for _, s1 := range csrIDs {
		if !slices.Contains(subjectIDs, s1) {
			return false
		}
	}

	// TODO(costin): this is a very basic check. We may allow the CSR to have O/OU matching a specific setting.
	if csr.Subject.CommonName != "" {
		return false
	}
	if len(csr.Subject.Organization) > 0 || len(csr.Subject.OrganizationalUnit) > 0 {
		return false
	}
	return true
}

// preSign : Validation checks to execute before signing certificates
// TODO(costin): this does not check everything... CN, Subject, other stuff...
func preSign(raOpts *IstioRAOptions, csrPEM []byte, subjectIDs []string, requestedLifetime time.Duration, forCA bool) (time.Duration, error) {
	if forCA {
		return requestedLifetime, raerror.NewError(raerror.CSRError,
			fmt.Errorf("unable to generate CA certifificates"))
	}
	if !ValidateCSR(csrPEM, subjectIDs) {
		return requestedLifetime, raerror.NewError(raerror.CSRError, fmt.Errorf(
			"unable to validate SAN Identities in CSR"))
	}
	// If the requested requestedLifetime is non-positive, apply the default TTL.
	// Shouldn't we get the lifetime from the cert ?
	lifetime := requestedLifetime
	if requestedLifetime.Seconds() <= 0 {
		lifetime = raOpts.DefaultCertTTL
	}
	// If the requested TTL is greater than maxCertTTL, return an error
	if requestedLifetime.Seconds() > raOpts.MaxCertTTL.Seconds() {
		return lifetime, raerror.NewError(raerror.TTLError, fmt.Errorf(
			"requested TTL %s is greater than the max allowed TTL %s", requestedLifetime, raOpts.MaxCertTTL))
	}
	return lifetime, nil
}
