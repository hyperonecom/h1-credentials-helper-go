package passport

import (
	"github.com/kuskoman/h1-credential-helper-go/auth/jwt"
	"github.com/kuskoman/h1-credential-helper-go/providers"
)

// AuthProvider is used to provide JWT auth using data from passport file
type AuthProvider struct {
	signer *jwt.RSA256Signer
}

// GetCredentialsHelper returns credential helper using passport file
func GetCredentialsHelper(location string) (providers.TokenAuthProvider, error) {
	var err error
	if location == "" {
		location, err = getDefaultPassportLocation()
	}

	if err != nil {
		return nil, err
	}

	passport, err := loadPassportFile(location)
	if err != nil {
		return nil, err
	}

	signer := &jwt.RSA256Signer{PrivateKey: passport.PrivateKey, KeyID: passport.CertificateID, Issuer: passport.Issuer, Subject: passport.SubjectID}

	provider := AuthProvider{signer: signer}

	return &provider, nil
}

// GetToken returns token used for signing requests
func (provider AuthProvider) GetToken(audience string) (string, error) {
	return provider.signer.GetJWT(audience)
}
