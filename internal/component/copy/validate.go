package copy

import (
	"errors"
	"os"
	"strings"
)

var (
	errFileNotExist = errors.New("file not exists")
)

// localFileValidate validate on exists file
func localFileValidate(filename string) error {
	if strings.TrimSpace(filename) == "" || filename == "" {
		return nil
	}

	_, err := os.Stat(filename)
	if err != nil {
		return errFileNotExist
	}

	return nil
}
