package file

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func Delete(path string) error {
	if Exists(path) {
		return os.Remove(path)
	}
	return fmt.Errorf("Dosya bulunamadı: %s", path)
}
