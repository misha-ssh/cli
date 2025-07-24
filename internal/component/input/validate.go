package input

import (
	"errors"
	"os"
	"strconv"
)

var (
	errPortIsNotString = errors.New("port is not string")
	errPortRange       = errors.New("port from 1 to 65535")
	errFileNotExist    = errors.New("file not exist")
)

func portValidate(s string) error {
	port, err := strconv.Atoi(s)
	if err != nil {
		return errPortIsNotString
	}

	if port < 1 || port > 65535 {
		return errPortRange
	}

	return nil
}

func fileExistsValidate(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return errFileNotExist
	}

	return nil
}

func aliasExistsValidate(alias string) error {
	return nil
}
