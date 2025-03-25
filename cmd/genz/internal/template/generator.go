package template

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/geekswamp/genz/internal/format"
)

func (m Make) Generate() error {
	path := filepath.Join(filepath.Clean(string(m.FilePath)), format.ToSnakeCase(string(m.FeatureName))+string(m.SuffixFile)+".go")
	dir := filepath.Dir(path)

	if _, err := os.Stat(path); err == nil {
		return errors.New("file already exists")
	}

	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s", dir)
	}

	file, err := os.Create(path)
	if err != nil {
		return errors.New("failed to create model")
	}

	defer file.Close()

	if err := m.Parse(file); err != nil {
		return fmt.Errorf("failed to execute template file %v", err)
	}

	return nil
}
