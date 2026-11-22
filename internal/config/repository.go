package config

import (
	"time"
	"uuid"
)

type Repo struct {
	Version      string    `json:"version"`
	RepositoryID string    `json:"repository_id"`
	CreatedAt    time.Time `json:"created_at"`
	Source       string    `json:"source"`
}

func NewRepo() *Repo {
	return &Repo{
		Version:      "1.0.0v",
		RepositoryID: uuid.NewV7().String(),
		CreatedAt:    time.Now(),
	}
}

func (r *Repo) SetSource(source string) {
	r.Source = source
}
