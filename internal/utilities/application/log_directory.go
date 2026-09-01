package utilities

import (
	"os"
	"path/filepath"
)

// LogDirectory returns the directory where logs are stored.
//
// If the directory does not exist, it is created.
//
// Returns:
//   - The directory where logs are stored.
//   - An error if the directory could not be created.
func LogDirectory() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	directory := filepath.Join(home, ".dns-updater", "logs")
	if err = os.MkdirAll(directory, 0o700); err != nil {
		return "", err
	}

	return directory, nil
}
