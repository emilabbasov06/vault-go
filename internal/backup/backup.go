package backup

import (
	"os"
	"path/filepath"
	"vault-go/internal/chunker"
	"vault-go/internal/config"
	"vault-go/internal/hasher"
	"vault-go/internal/objectstore"
	"vault-go/internal/scanner"
)

func BackupProcess(bFolder string) error {
	repo, err := config.ReadConf()
	if err != nil {
		return err
	}

	repo.Source = bFolder
	if err = config.WriteConf(repo); err != nil {
		return err
	}

	homeDir, _ := config.GetHomeDir()
	path := filepath.Join(homeDir, bFolder)
	entries, err := scanner.Scan(path)
	if err != nil {
		return err
	}

	store := objectstore.NewObjectStore(bFolder)
	for _, entry := range entries {
		if entry.IsDir {
			continue
		}

		file, err := os.Open(entry.Path)
		if err != nil {
			return err
		}

		if err = chunker.ChunkProcess(file, 1024*1024, func(chunk *chunker.Chunk) error {
			if err = store.Put(hasher.Hash(chunk.Data), chunk.Data); err != nil {
				return err
			}
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
