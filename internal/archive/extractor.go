package archive

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
)

func Extract(inputFilePath, outputDir string) error {
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	gzipReader, err := gzip.NewReader(inputFile)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		targetPath := filepath.Join(outputDir, header.Name)
		info := header.FileInfo()

		if info.IsDir() {
			if err := os.MkdirAll(targetPath, info.Mode()); err != nil {
				return err
			}
			continue
		}

		file, err := os.OpenFile(targetPath, os.O_CREATE|os.O_RDWR, info.Mode())
		if err != nil {
			return err
		}

		if _, err := io.Copy(file, tarReader); err != nil {
			file.Close()
			return err
		}
		file.Close()
	}

	return nil
}
