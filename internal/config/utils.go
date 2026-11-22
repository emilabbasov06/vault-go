package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetHomeDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	return homeDir, nil
}

func GetMainDir() (string, error) {
	homeDir, err := GetHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, "Documents", "VaultGoBackup", ".vaultgo"), nil
}

func MainSetup(mainDir string) error {
	subDirs := []string{
		filepath.Join(mainDir, "objects"),
		filepath.Join(mainDir, "snapshots"),
		filepath.Join(mainDir, "tmp"),
	}

	for _, subDir := range subDirs {
		if err := os.MkdirAll(subDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	return nil
}
