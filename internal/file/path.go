package file

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsolutePath(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func JoinPaths(basePath, subPath string) string {
	return filepath.Join(basePath, subPath)
}

func ValidateDirectory(path string) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}
	if !fileInfo.IsDir() {
		return fmt.Errorf("%s bir dizin değil", path)
	}
	return nil
}
