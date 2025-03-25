package mod

import (
	"errors"
	"os"

	"github.com/geekswamp/zen/pkg/file"
	"golang.org/x/mod/modfile"
)

var mod = "go.mod"

func IsExist() bool {
	_, exist := file.IsExist(mod)
	return exist
}

func GetModuleName() (*string, error) {
	if !IsExist() {
		return nil, errors.New("go module not found. Please ensure that this is a Go project and a go.mod file exists in the root directory")
	}

	data, err := os.ReadFile(mod)
	if err != nil {
		return nil, err
	}

	modFile, err := modfile.Parse(mod, data, nil)
	if err != nil {
		return nil, err
	}

	if modFile.Module == nil {
		return nil, err
	}

	return &modFile.Module.Mod.Path, nil
}
