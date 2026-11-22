package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func confSetup(mainDir string, repo *Repo) error {
	data, err := json.Marshal(repo)
	if err != nil {
		return fmt.Errorf("coudn't convert struct to json data: %w", err)
	}

	path := filepath.Join(mainDir, "config.json")
	if err = os.WriteFile(path, data, 0755); err != nil {
		return fmt.Errorf("coudn't create 'config.json' file for configuration: %w", err)
	}

	return nil
}

func FolderSetup() (*Repo, error) {
	mainDir, err := GetMainDir()
	if err != nil {
		return nil, err
	}

	if err = MainSetup(mainDir); err != nil {
		return nil, err
	}

	repo := NewRepo()
	if err = confSetup(mainDir, repo); err != nil {
		return nil, err
	}

	return repo, nil
}

func GetConf() (string, error) {
	mainDir, err := GetMainDir()
	if err != nil {
		return "", err
	}

	path := filepath.Join(mainDir, "config.json")
	configData, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("couldn't read 'config.json': %w", err)
	}

	repo := NewRepo()
	if err = json.Unmarshal(configData, repo); err != nil {
		return "", fmt.Errorf("couldn't conver 'config.json' to real struct object: %w", err)
	}

	template := `
VaultGo Backup

Repository ID: %s
Version:       %s
Created At:    %s
Source:      %s
`
	return fmt.Sprintf(template, repo.RepositoryID, repo.Version, repo.CreatedAt, repo.Source), nil
}
