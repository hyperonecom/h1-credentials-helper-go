package h1

import (
	"github.com/kuskoman/h1-credentials-helper-go/providers"
	"github.com/kuskoman/h1-credentials-helper-go/providers/passport"
)

// GetPassportCredentialsHelper creates credentials helper using HyperOne passport file
func GetPassportCredentialsHelper(location string) (providers.TokenAuthProvider, error) {
	provider, err := passport.GetCredentialsHelper(location)
	if err != nil {
		return nil, err
	}

	return provider, nil
}
