package test

import (
	"errors"
	"path/filepath"
	"runtime"
)

func GetCurrentDir() (string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return "", errors.New("failed to get runtime.Caller")
	}
	return filepath.Abs(filepath.Dir(filename))
}
