package functions

import (
	"os"
	"path/filepath"
)

type FileFuncs struct{}

func (FileFuncs) Exists(inputPath string) bool {
	_, err := os.Stat(inputPath)

	return err == nil
}

func (FileFuncs) IsDir(inputPath string) bool {
	info, err := os.Stat(inputPath)

	return err == nil && info.IsDir()
}

func (FileFuncs) IsFile(inputPath string) bool {
	info, err := os.Stat(inputPath)

	return err == nil && !info.IsDir()
}

func (FileFuncs) Read(inputPath string) (string, error) {
	bytes, err := os.ReadFile(filepath.Clean(inputPath))

	return string(bytes), err
}

func (FileFuncs) ReadDir(inputPath string) ([]string, error) {
	entries, err := os.ReadDir(filepath.Clean(inputPath))
	if err != nil {
		return nil, err
	}

	var items []string
	for _, file := range entries {
		items = append(items, file.Name())
	}

	return items, nil
}

func (FileFuncs) Stat(inputPath string) (os.FileInfo, error) {
	return os.Stat(inputPath)
}

func (FileFuncs) Write(outputPath string, data any) error {
	return os.WriteFile(filepath.Clean(outputPath), []byte(ConversionFuncs{}.ToString(data)), 0o600)
}
