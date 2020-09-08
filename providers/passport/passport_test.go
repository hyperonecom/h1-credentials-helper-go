package passport

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGettingToken(t *testing.T) {
	provider, err := GetCredentialsHelper("fixtures/validPassport.json")
	require.NoError(t, err)

	_, err = provider.GetToken("exampleAudience")
	require.NoError(t, err)
}
