package list

import (
	"errors"
	"github.com/charmbracelet/huh"
	"github.com/misha-ssh/kernel/pkg/connect"
)

var (
	errNotFoundConnections = errors.New("not found connections")
)

func Run(connections *connect.Connections) (*Fields, error) {
	var aliases []string
	var fields Fields

	if len(connections.Connects) == 0 {
		return nil, errNotFoundConnections
	}

	for _, conn := range connections.Connects {
		aliases = append(aliases, conn.Alias)
	}

	options := huh.NewOptions(aliases...)

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Connection").
				Description("test description").
				Options(options...).
				Value(&fields.Alias),
		)).WithShowHelp(true).Run()

	if err != nil {
		return nil, err
	}

	return &fields, nil
}
