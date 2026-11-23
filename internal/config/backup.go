package config

import (
	"fmt"
	"path/filepath"
	"vault-go/internal/scanner"
)

func BackupProcess(bFolder string) error {
	repo, err := ReadConf()
	if err != nil {
		return err
	}

	repo.Source = bFolder
	if err = WriteConf(repo); err != nil {
		return err
	}

	homeDir, _ := GetHomeDir()
	path := filepath.Join(homeDir, bFolder)
	entries, err := scanner.Scan(path)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		fmt.Println(entry)
	}

	return nil
}
