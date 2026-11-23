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
