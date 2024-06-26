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

// Provides utility methods to generate X.509 certificates with different
// options. This implementation is Largely inspired from
// https://golang.org/src/crypto/tls/generate_cert.go.

package util

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
	"sync"
	"time"
)

// KeyCertBundle stores the cert, private key, cert chain and root cert for an entity. It is thread safe.
// The cert and privKey should be a public/private key pair.
// The cert should be verifiable from the rootCert through the certChain.
// cert and priveKey are pointers to the cert/key parsed from certBytes/privKeyBytes.
type KeyCertBundle struct {
	// certBytes holds the leaf certificate - associated with the private key.
	certBytes    []byte
	cert         *x509.Certificate
	privKeyBytes []byte
	privKey      *crypto.PrivateKey

	// certChainBytes holds the intermediary certificates. Empty for the auto-generated root.
	certChainBytes []byte

	// rootCertBytes has the PEM concatenated list of roots, loaded from secret or file.
	rootCertBytes []byte
	// mutex protects the R/W to all keys and certs.
	mutex sync.RWMutex
}

// NewKeyCertBundleFromPem returns a new KeyCertBundle, regardless of whether or not the key can be correctly parsed.
func NewKeyCertBundleFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) *KeyCertBundle {
	bundle := &KeyCertBundle{}
	bundle.setAllFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes)
	return bundle
}

// SplitTlsCrt will break a tls.crt PEM containing leaf plus intermediaries into
// 2 PEMs, one containing the leaf and the other the intermediaries, as is expected
// by the Istiod internals, verifications, etc.
//
// If the input contains only one certificate - it will be returned unchanged.
func SplitTlsCrt(tlsCrt []byte) ([]byte, []byte) {
	p, rest := pem.Decode(tlsCrt)
	if p == nil || rest == nil {
		return tlsCrt, nil
	}
	return pem.EncodeToMemory(p), rest
}

// NewVerifiedKeyCertBundleFromPem returns a new KeyCertBundle, or error if the provided certs failed the
// verification.
func NewVerifiedKeyCertBundleFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) (
	*KeyCertBundle, error,
) {
	bundle := &KeyCertBundle{}
	if err := bundle.VerifyAndSetAll(certBytes, privKeyBytes, certChainBytes, rootCertBytes); err != nil {
		return nil, err
	}
	return bundle, nil
}

// NewVerifiedKeyCertBundleFromFile returns a new KeyCertBundle, or error if the provided certs failed the
// verification.
func NewVerifiedKeyCertBundleFromFile(certFile string, privKeyFile string, certChainFiles []string, rootCertFile string) (
	*KeyCertBundle, error,
) {
	certBytes, err := os.ReadFile(certFile)
	if err != nil {
		return nil, err
	}
	chaincerts, _ := SplitPemEncodedCertificates(certBytes)
	intNames := []string{}
	for _, r := range chaincerts {
		intNames = append(intNames, r.Subject.String())
	}

	// If the cert file is tls.key - it holds both leaf and intermediaries.
	leafBytes, certChainBytes := SplitTlsCrt(certBytes)
	certBytes = leafBytes

	privKeyBytes, err := os.ReadFile(privKeyFile)
	if err != nil {
		return nil, err
	}

	// Standalone, separate certChainFiles
	if len(certChainFiles) > 0 {
		for _, f := range certChainFiles {
			var b []byte

			if b, err = os.ReadFile(f); err != nil {
				return nil, err
			}

			certChainBytes = append(certChainBytes, b...)
		}
	}
	rootCertBytes, err := os.ReadFile(rootCertFile)
	if err != nil {
		return nil, err
	}
	return NewVerifiedKeyCertBundleFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes)
}

// NewKeyCertBundleWithRootCertFromFile returns a new KeyCertBundle with the root cert without verification.
func NewKeyCertBundleWithRootCertFromFile(rootCertFile string) (*KeyCertBundle, error) {
	var rootCertBytes []byte
	var err error
	if rootCertFile == "" {
		rootCertBytes = []byte{}
	} else {
		rootCertBytes, err = os.ReadFile(rootCertFile)
		if err != nil {
			return nil, err
		}
	}
	return &KeyCertBundle{
		certBytes:      []byte{},
		cert:           nil,
		privKeyBytes:   []byte{},
		privKey:        nil,
		certChainBytes: []byte{},
		rootCertBytes:  rootCertBytes,
	}, nil
}

// GetAllPem returns all key/cert PEMs in KeyCertBundle together. Getting all values together avoids inconsistency.
func (b *KeyCertBundle) GetAllPem() (certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) {
	b.mutex.RLock()
	certBytes = copyBytes(b.certBytes)
	privKeyBytes = copyBytes(b.privKeyBytes)
	certChainBytes = copyBytes(b.certChainBytes)
	rootCertBytes = copyBytes(b.rootCertBytes)
	b.mutex.RUnlock()
	return
}

// GetAll returns all key/cert in KeyCertBundle together. Getting all values together avoids inconsistency.
// NOTE: Callers should not modify the content of cert and privKey.
func (b *KeyCertBundle) GetAll() (cert *x509.Certificate, privKey *crypto.PrivateKey, certChainBytes,
	rootCertBytes []byte,
) {
	b.mutex.RLock()
	cert = b.cert
	privKey = b.privKey
	certChainBytes = copyBytes(b.certChainBytes)
	rootCertBytes = copyBytes(b.rootCertBytes)
	b.mutex.RUnlock()
	return
}

// GetCertChainPem returns the certificate chain PEM.
func (b *KeyCertBundle) GetCertChainPem() []byte {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return copyBytes(b.certChainBytes)
}

// GetRootCertPem returns the root certificate PEM.
func (b *KeyCertBundle) GetRootCertPem() []byte {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	return copyBytes(b.rootCertBytes)
}

// VerifyAndSetAll verifies the key/certs, and sets all key/certs in KeyCertBundle together.
// Setting all values together avoids inconsistency.
//
// This method expects 'certBytes' to include a single cert - corresponding to the
// private key.
// The intermediary certificates (if any) must be included in the certChainBytes.
// RootCert and certChainBytes are concatenated PEM certs.
func (b *KeyCertBundle) VerifyAndSetAll(certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) error {
	if err := Verify(certBytes, privKeyBytes, certChainBytes, rootCertBytes); err != nil {
		return err
	}
	b.setAllFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes)
	return nil
}

// Setting all values together avoids inconsistency.
func (b *KeyCertBundle) setAllFromPem(certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) {
	b.mutex.Lock()
	b.certBytes = copyBytes(certBytes)
	b.privKeyBytes = copyBytes(privKeyBytes)
	b.certChainBytes = copyBytes(certChainBytes)
	b.rootCertBytes = copyBytes(rootCertBytes)
	// cert and privKey are always reset to point to new addresses. This avoids modifying the pointed structs that
	// could be still used outside of the class.
	b.cert, _ = ParsePemEncodedCertificate(certBytes)
	privKey, _ := ParsePemEncodedKey(privKeyBytes)
	b.privKey = &privKey
	b.mutex.Unlock()
}

// CertOptions returns the certificate config based on currently stored cert.
func (b *KeyCertBundle) CertOptions() (*CertOptions, error) {
	b.mutex.RLock()
	defer b.mutex.RUnlock()
	ids, err := ExtractIDs(b.cert.Extensions)
	if err != nil {
		return nil, fmt.Errorf("failed to extract id %v", err)
	}
	if len(ids) != 1 {
		return nil, fmt.Errorf("expect single id from the cert, found %v", ids)
	}

	opts := &CertOptions{
		Host:      ids[0],
		Org:       b.cert.Issuer.Organization[0],
		IsCA:      b.cert.IsCA,
		TTL:       b.cert.NotAfter.Sub(b.cert.NotBefore),
		IsDualUse: ids[0] == b.cert.Subject.CommonName,
	}

	switch (*b.privKey).(type) {
	case *rsa.PrivateKey:
		size, err := GetRSAKeySize(*b.privKey)
		if err != nil {
			return nil, fmt.Errorf("failed to get RSA key size: %v", err)
		}
		opts.RSAKeySize = size
	case *ecdsa.PrivateKey:
		opts.ECSigAlg = EcdsaSigAlg
	default:
		return nil, errors.New("unknown private key type")
	}

	return opts, nil
}

// UpdateVerifiedKeyCertBundleFromFile Verifies and updates KeyCertBundle with new certs
func (b *KeyCertBundle) UpdateVerifiedKeyCertBundleFromFile(certFile string, privKeyFile string, certChainFiles []string, rootCertFile string) error {
	certBytes, err := os.ReadFile(certFile)
	if err != nil {
		return err
	}
	privKeyBytes, err := os.ReadFile(privKeyFile)
	if err != nil {
		return err
	}
	certChainBytes := []byte{}
	if len(certChainFiles) != 0 {
		for _, f := range certChainFiles {
			var b []byte
			if b, err = os.ReadFile(f); err != nil {
				return err
			}

			certChainBytes = append(certChainBytes, b...)
		}
	}
	rootCertBytes, err := os.ReadFile(rootCertFile)
	if err != nil {
		return err
	}

	err = b.VerifyAndSetAll(certBytes, privKeyBytes, certChainBytes, rootCertBytes)
	if err != nil {
		return err
	}

	return nil
}

// ExtractRootCertExpiryTimestamp returns the unix timestamp when the root becomes expires.
func (b *KeyCertBundle) ExtractRootCertExpiryTimestamp() (float64, error) {
	return extractCertExpiryTimestamp("root cert", b.GetRootCertPem())
}

// ExtractCACertExpiryTimestamp returns the unix timestamp when the cert chain becomes expires.
func (b *KeyCertBundle) ExtractCACertExpiryTimestamp() (float64, error) {
	return extractCertExpiryTimestamp("CA cert", b.GetCertChainPem())
}

// TimeBeforeCertExpires returns the time duration before the cert gets expired.
// It returns an error if it failed to extract the cert expiration timestamp.
// The returned time duration could be a negative value indicating the cert has already been expired.
func TimeBeforeCertExpires(certBytes []byte, now time.Time) (time.Duration, error) {
	if len(certBytes) == 0 {
		return 0, fmt.Errorf("no certificate found")
	}

	certExpiryTimestamp, err := extractCertExpiryTimestamp("cert", certBytes)
	if err != nil {
		return 0, fmt.Errorf("failed to extract cert expiration timestamp: %v", err)
	}

	certExpiry := time.Duration(certExpiryTimestamp-float64(now.Unix())) * time.Second
	return certExpiry, nil
}

// Verify that the cert chain, root cert and key/cert match.
func Verify(certBytes, privKeyBytes, certChainBytes, rootCertBytes []byte) error {
	// Verify the cert can be verified from the root cert through the cert chain.
	rcp := x509.NewCertPool()
	rcp.AppendCertsFromPEM(rootCertBytes)

	icp := x509.NewCertPool()
	icp.AppendCertsFromPEM(certChainBytes)

	opts := x509.VerifyOptions{
		Intermediates: icp,
		Roots:         rcp,
	}
	cert, err := ParsePemEncodedCertificate(certBytes)
	if err != nil {
		return fmt.Errorf("failed to parse cert PEM: %v", err)
	}
	chains, err := cert.Verify(opts)

	if len(chains) == 0 || err != nil {
		return fmt.Errorf(
			"cannot verify the cert with the provided root chain and cert "+
				"pool with error: %v", err)
	}

	// Verify that the key can be correctly parsed.
	if _, err = ParsePemEncodedKey(privKeyBytes); err != nil {
		return fmt.Errorf("failed to parse private key PEM: %v", err)
	}

	// Verify the cert and key match.
	if _, err := tls.X509KeyPair(certBytes, privKeyBytes); err != nil {
		return fmt.Errorf("the cert does not match the key: %v", err)
	}

	return nil
}

func extractCertExpiryTimestamp(certType string, certPem []byte) (float64, error) {
	cert, err := ParsePemEncodedCertificate(certPem)
	if err != nil {
		return -1, fmt.Errorf("failed to parse the %s: %v", certType, err)
	}

	end := cert.NotAfter
	expiryTimestamp := float64(end.Unix())
	if end.Before(time.Now()) {
		return expiryTimestamp, fmt.Errorf("expired %s found, x509.NotAfter %v, please transit your %s", certType, end, certType)
	}
	return expiryTimestamp, nil
}

func copyBytes(src []byte) []byte {
	bs := make([]byte, len(src))
	copy(bs, src)
	return bs
}
