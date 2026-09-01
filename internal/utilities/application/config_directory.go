package utilities

import (
	"os"
	"path/filepath"
)

// ConfigDirectory returns the directory where the config file is stored.
//
// If the directory does not exist, it is created.
//
// Returns:
//   - The directory where the config file is stored.
//   - An error if the directory could not be created.
func ConfigDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	directory := filepath.Join(home, ".config", "dns-updater")
	if err = os.MkdirAll(directory, 0o700); err != nil {
		return "", err
	}

	return directory, nil
}
