package input

import (
	"errors"
	"os"
	"strconv"

	"github.com/misha-ssh/kernel/pkg/connect"
	"github.com/misha-ssh/kernel/pkg/kernel"
)

const (
	MinPort = 1
	MaxPort = 65535
)

var (
	errPortIsNotString = errors.New("port is not string")
	errPortRange       = errors.New("port from 1 to 65535")
	errFileNotExist    = errors.New("file not exist")
	errGetConnections  = errors.New("get connections error")
	errAliasExists     = errors.New("alias exists")

	preloadedConnections    *connect.Connections
	errPreloadedConnections error
)

func init() {
	connections, err := kernel.List()
	if err != nil {
		errPreloadedConnections = err
	}

	preloadedConnections = connections
}

func aliasExistsValidate(alias string) error {
	if errPreloadedConnections != nil {
		return errGetConnections
	}

	for _, connection := range preloadedConnections.Connects {
		if connection.Alias == alias {
			return errAliasExists
		}
	}

	return nil
}

func portValidate(s string) error {
	port, err := strconv.Atoi(s)
	if err != nil {
		return errPortIsNotString
	}

	if port < MinPort || port > MaxPort {
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
