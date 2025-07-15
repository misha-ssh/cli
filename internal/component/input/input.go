package input

import (
	"errors"
	"reflect"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	fieldNumberAlias = iota
	fieldNumberLogin
	fieldNumberPassword
	fieldNumberPort
	fieldNumberPrivateKey
)

var (
	errInitInput        = errors.New("input initialization failed")
	errInvalidModelType = errors.New("invalid model type")
)

type Fields struct {
	Alias      string
	Login      string
	Password   string
	Port       int
	PrivateKey string
}

func (f *Fields) count() int {
	fieldsType := reflect.TypeOf(*f)
	return fieldsType.NumField()
}

func (f *Fields) getNameByNumber(number int) string {
	val := reflect.ValueOf(f)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	typ := val.Type()

	if number > typ.NumField() {
		return ""
	}

	return typ.Field(number).Name
}

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
		return nil, err
	}

	return &Fields{
		Alias:      currentModel.inputs[fieldNumberAlias].Value(),
		Login:      currentModel.inputs[fieldNumberLogin].Value(),
		Password:   currentModel.inputs[fieldNumberPassword].Value(),
		Port:       port,
		PrivateKey: currentModel.inputs[fieldNumberPrivateKey].Value(),
	}, nil
}
