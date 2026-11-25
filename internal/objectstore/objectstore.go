package objectstore

import (
	"fmt"
	"os"
	"path/filepath"
	"vault-go/internal/config"
)

type ObjectStore struct {
	Root string
}

func NewObjectStore(root string) *ObjectStore {
	return &ObjectStore{Root: root}
}

func (o *ObjectStore) Put(hash string, data []byte) error {
	if hash == "" || data == nil {
		return fmt.Errorf("invalid params, can't store the object")
	}

	mainDir, err := config.GetMainDir()
	if err != nil {
		return err
	}

	objectFolder := hash[:2]
	object := hash[2:]
	objects := filepath.Join(mainDir, "objects", objectFolder)

	fmt.Println(objectFolder, object, mainDir)

	if err = os.MkdirAll(objects, 0755); err != nil {
		return err
	}

	file := filepath.Join(objects, object)
	if err = os.WriteFile(file, []byte(""), 0755); err != nil {
		return err
	}

	return nil
}
