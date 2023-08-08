package file

import (
	"os"
	"path/filepath"
)

func CreateDirectory(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

func ListFiles(path string) ([]string, error) {
	var files []string

	err := filepath.Walk(path, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, filePath)
		}
		return nil
	})

	return files, err
}
