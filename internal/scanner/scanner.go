package scanner

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func Scan(source string) ([]Entry, error) {
	entries := []Entry{}
	if err := filepath.WalkDir(source, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		entry := NewEntry()
		entry.Path = path
		entry.IsDir = d.IsDir()
		entry.ModTime = info.ModTime()
		entry.Mode = info.Mode()
		entry.Size = info.Size()

		entries = append(entries, *entry)

		return nil
	}); err != nil {
		return nil, fmt.Errorf("couldn't read backup directory: %w", err)
	}

	return entries, nil
}
