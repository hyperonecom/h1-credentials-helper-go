package credentials

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatingPassportProvider(t *testing.T) {
	_, err := GetPassportCredentialsHelper("providers/passport/fixtures/validPassport.json")
	require.NoError(t, err)
	_, err = GetPassportCredentialsHelper("providers/passport/fixtures/invalidPassport.json")
	require.Error(t, err)
}
