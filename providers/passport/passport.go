package passport

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/kuskoman/h1-credential-helper-go/providers"
)

// Passport contains data used for authenticating user or service account
type Passport struct {
	SubjectID     string `json:"subject_id"`
	CertificateID string `json:"certificate_id"`
	Issuer        string `json:"issuer"`
	PrivateKey    string `json:"private_key"`
	PublicKey     string `json:"public_key"`
}

// GetPassportCredentialHelper returns credential helper using
func GetPassportCredentialHelper(location string) (*providers.JWTAuthProvider, error) {
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

	return passport, nil
}

func loadPassportFile(location string) (*Passport, error) {
	file, err := os.Open(location)
	if err != nil {
		return nil, fmt.Errorf("failed to open passport file: %w", err)
	}

	defer func() { _ = file.Close() }()

	var passport Passport
	err = json.NewDecoder(file).Decode(&passport)
	if err != nil {
		return nil, fmt.Errorf("failed to parse passport file: %w", err)
	}

	err = passport.validate()
	if err != nil {
		return nil, fmt.Errorf("passport file validation failed: %w", err)
	}

	return &passport, nil
}

func (passport *Passport) validate() error {
	if passport.Issuer == "" {
		return errors.New("issuer is empty")
	}

	if passport.CertificateID == "" {
		return errors.New("certificate ID is empty")
	}

	if passport.PrivateKey == "" {
		return errors.New("private key is missing")
	}

	if passport.SubjectID == "" {
		return errors.New("subject is empty")
	}

	return nil
}
