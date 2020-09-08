package passport

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

// CredentialsFile contains data used for authenticating user or service account
type CredentialsFile struct {
	SubjectID     string `json:"subject_id"`
	CertificateID string `json:"certificate_id"`
	Issuer        string `json:"issuer"`
	PrivateKey    string `json:"private_key"`
	PublicKey     string `json:"public_key"`
}

func loadPassportFile(location string) (*CredentialsFile, error) {
	file, err := os.Open(location)
	if err != nil {
		return nil, fmt.Errorf("failed to open passport file: %w", err)
	}

	defer func() { _ = file.Close() }()

	var passport CredentialsFile
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

func (passportFile *CredentialsFile) validate() error {
	if passportFile.Issuer == "" {
		return errors.New("issuer is empty")
	}

	if passportFile.CertificateID == "" {
		return errors.New("certificate ID is empty")
	}

	if passportFile.PrivateKey == "" {
		return errors.New("private key is missing")
	}

	if passportFile.SubjectID == "" {
		return errors.New("subject is empty")
	}

	return nil
}

func getDefaultPassportLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".h1", "passport.json"), nil
}
