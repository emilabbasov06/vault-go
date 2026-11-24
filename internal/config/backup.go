package config

import (
	"fmt"
	"os"
	"path/filepath"
	"vault-go/internal/chunker"
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
		if entry.IsDir {
			continue
		}

		file, err := os.Open(entry.Path)
		if err != nil {
			return err
		}

		if err = chunker.ChunkProcess(file, 600*1024, func(chunk *chunker.Chunk) error {
			fmt.Fprintf(os.Stdout, "file=%s, chunk_size=%d\n", entry.Path, len(chunk.Data))
			return nil
		}); err != nil {
			return err
		}

		if err = file.Close(); err != nil {
			return err
		}
	}

	return nil
}
