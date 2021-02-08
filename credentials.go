package credentials

import (
	"github.com/hyperonecom/h1-credentials-helper-go/providers"
	"github.com/hyperonecom/h1-credentials-helper-go/providers/passport"
)

// GetPassportCredentialsHelper creates credentials helper using HyperOne passport file
func GetPassportCredentialsHelper(location string) (providers.TokenAuthProvider, error) {
	return passport.GetCredentialsHelper(location)
}
