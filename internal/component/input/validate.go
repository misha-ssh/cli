package input

import (
	"errors"
	"strconv"
)

var (
	errPortIsNotString = errors.New("port is not string")
	errPortRange       = errors.New("port from 1 to 65535")
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
