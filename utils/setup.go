package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func BackupFolderSetup() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get user home directory: %w", err)
	}

	mainDir := filepath.Join(homeDir, "Documents", "VaultGoBackup", ".vaultgo")

	subDirs := []string{
		filepath.Join(mainDir, "objects"),
		filepath.Join(mainDir, "snapshots"),
		filepath.Join(mainDir, "tmp"),
	}

	for _, subDir := range subDirs {
		if err = os.MkdirAll(subDir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}

	config := filepath.Join(mainDir, "config.json")
	configTemplate := []byte(`// Simple config.json file for VaultGo setup.`)
	if err = os.WriteFile(config, configTemplate, 0755); err != nil {
		return fmt.Errorf("failed to create 'config.json' file: %w", err)
	}

	return nil
}
