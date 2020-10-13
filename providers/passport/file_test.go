package passport

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadPassportFile(t *testing.T) {
	_, err := loadPassportFile("fixtures/validPassport.json")
	require.NoError(t, err)
}

func TestLoadPassportFile_invalid(t *testing.T) {
	passport, err := loadPassportFile("fixtures/invalidPassport.json")
	require.EqualError(t, err, "passport file validation failed: private key is missing")

	assert.Nil(t, passport)
}
