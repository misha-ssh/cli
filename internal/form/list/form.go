package list

import (
	"github.com/charmbracelet/huh"
	"github.com/misha-ssh/kernel/pkg/connect"
)

func Run(connections *connect.Connections) (*Fields, error) {
	var aliases []string
	var fields Fields

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
