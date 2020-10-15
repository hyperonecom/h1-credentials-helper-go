package credentials

import (
	"github.com/kuskoman/h1-credentials-helper-go/providers"
	"github.com/kuskoman/h1-credentials-helper-go/providers/passport"
)

// GetPassportCredentialsHelper creates credentials helper using HyperOne passport file
func GetPassportCredentialsHelper(location string) (providers.TokenAuthProvider, error) {
	return passport.GetCredentialsHelper(location)
}
