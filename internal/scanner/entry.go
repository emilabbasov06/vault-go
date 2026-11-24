package scanner

import (
	"io/fs"
	"time"
)

type Entry struct {
	Path    string
	IsDir   bool
	Size    int64
	ModTime time.Time
	Mode    fs.FileMode
}

func NewEntry() *Entry {
	return &Entry{}
}

func (e *Entry) SetEntry(path string, isDir bool, size int64, modTime time.Time, mode fs.FileMode) {
	e.Path = path
	e.IsDir = isDir
	e.Size = size
	e.ModTime = modTime
	e.Mode = mode
}
