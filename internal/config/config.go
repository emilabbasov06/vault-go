package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
	"uuid"
)

type VGORepoInfo struct {
	Version      string    `json:"version"`
	RepositoryID string    `json:"repository_id"`
	CreatedAt    time.Time `json:"created_at"`
	Location     string    `json:"location"`
	Status       string    `json:"status"`
}

var repoInfo = VGORepoInfo{
	Version:      "1.10.1",
	RepositoryID: uuid.NewV7().String(),
	CreatedAt:    time.Now(),
	Status:       "active",
}

func configFileSetup(mainDir string) error {
	config := filepath.Join(mainDir, "config.json")
	configData, err := json.Marshal(repoInfo)
	if err != nil {
		return err
	}

	if err = os.WriteFile(config, configData, 0755); err != nil {
		return err
	}

	return nil
}

func BackupFolderSetup() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get user home directory: %w", err)
	}

	mainDir := filepath.Join(homeDir, "Documents", "VaultGoBackup", ".vaultgo")
	repoInfo.Location = mainDir

	subDirs := []string{
		filepath.Join(mainDir, "objects"),
		filepath.Join(mainDir, "snapshots"),
		filepath.Join(mainDir, "tmp"),
	}

	for _, subDir := range subDirs {
		if err = os.MkdirAll(subDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create directory: %w", err)
		}
	}

	if err = configFileSetup(mainDir); err != nil {
		return "", fmt.Errorf("failed to create 'config.json' file: %w", err)
	}

	return repoInfo.Version, nil
}

func GetRepoInfo() string {
	template := `
VaultGo Backup

Repository ID: %s
Version:       %s
Created At:    %s
Location:      %s
Status:        %s
`
	return fmt.Sprintf(template, repoInfo.RepositoryID, repoInfo.Version, repoInfo.CreatedAt, repoInfo.Location, repoInfo.Status)
}
