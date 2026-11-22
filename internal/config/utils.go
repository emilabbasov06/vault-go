package config

import (
	"encoding/json"
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

func ReadConf() (*Repo, error) {
	mainDir, err := GetMainDir()
	if err != nil {
		return nil, err
	}

	path := filepath.Join(mainDir, "config.json")
	configData, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("couldn't read 'config.json': %w", err)
	}

	repo := NewRepo()
	if err = json.Unmarshal(configData, repo); err != nil {
		return nil, fmt.Errorf("couldn't conver 'config.json' to real struct object: %w", err)
	}

	return repo, nil
}

func WriteConf(repo *Repo) error {
	data, err := json.Marshal(repo)
	if err != nil {
		return fmt.Errorf("coudn't convert struct to json data: %w", err)
	}

	mainDir, err := GetMainDir()
	if err != nil {
		return err
	}

	path := filepath.Join(mainDir, "config.json")
	if err = os.WriteFile(path, data, 0755); err != nil {
		return fmt.Errorf("coudn't create 'config.json' file for configuration: %w", err)
	}

	return nil
}
