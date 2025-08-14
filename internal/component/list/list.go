package list

import (
	"errors"

	"github.com/charmbracelet/huh"
	"github.com/misha-ssh/kernel/pkg/connect"
)

var (
	errNotFoundConnections = errors.New("not found connections")
)

// Run Get selected connection from list
func Run(connections *connect.Connections) (*connect.Connect, error) {
	var aliases []string
	var selectedAlias string

	if len(connections.Connects) == 0 {
		return nil, errNotFoundConnections
	}

	aliasToConn := make(map[string]*connect.Connect)

	for _, conn := range connections.Connects {
		aliases = append(aliases, conn.Alias)
		aliasToConn[conn.Alias] = &conn
	}

	err := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Select Connection").
				Options(huh.NewOptions(aliases...)...).
				Value(&selectedAlias),
		)).WithShowHelp(true).Run()
	if err != nil {
		return nil, err
	}

	selectedConn, exists := aliasToConn[selectedAlias]
	if !exists {
		return nil, errNotFoundConnections
	}

	return selectedConn, nil
}
