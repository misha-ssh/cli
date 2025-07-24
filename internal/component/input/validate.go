package input

import (
	"errors"
	"os"
	"strconv"
	"strings"

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
	errFileNotExist    = errors.New("file not exists")
	errGetConnections  = errors.New("get connections error")
	errAliasExists     = errors.New("alias exists")
	errAliasIsNotEmpty = errors.New("alias is not empty")

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

func aliasValidate(alias string) error {
	if strings.TrimSpace(alias) == "" || alias == "" {
		return errAliasIsNotEmpty
	}

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

func fileValidate(filename string) error {
	_, err := os.Stat(filename)
	if err != nil {
		return errFileNotExist
	}

	return nil
}
