package input

import (
	"errors"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

var (
	errInitInput        = errors.New("input initialization failed")
	errInvalidModelType = errors.New("invalid model type")
	errFormatPortNumber = errors.New("invalid port number")
)

func Run() (*Fields, error) {
	returnedModel, err := tea.NewProgram(initModel()).Run()
	if err != nil {
		return nil, errInitInput
	}

	currentModel, ok := returnedModel.(model)
	if !ok {
		return nil, errInvalidModelType
	}

	port, err := strconv.Atoi(currentModel.inputs[fieldNumberPort].Value())
	if err != nil {
		return nil, errFormatPortNumber
	}

	return &Fields{
		Alias:      currentModel.inputs[fieldNumberAlias].Value(),
		Login:      currentModel.inputs[fieldNumberLogin].Value(),
		Password:   currentModel.inputs[fieldNumberPassword].Value(),
		Port:       port,
		PrivateKey: currentModel.inputs[fieldNumberPrivateKey].Value(),
	}, nil
}
