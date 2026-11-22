package config

import (
	"fmt"
)

func FolderSetup() (*Repo, error) {
	mainDir, err := GetMainDir()
	if err != nil {
		return nil, err
	}

	if err = MainSetup(mainDir); err != nil {
		return nil, err
	}

	repo := NewRepo()
	if err = WriteConf(repo); err != nil {
		return nil, err
	}

	return repo, nil
}

func GetConf() (string, error) {
	repo, err := ReadConf()
	if err != nil {
		return "", err
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
