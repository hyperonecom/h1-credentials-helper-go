package passport

import (
	"fmt"
	"os"
	"path/filepath"
)

func getDefaultPassportLocation() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return filepath.Join(homeDir, ".h1", "passport.json"), nil
}
