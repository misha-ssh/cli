package input

import (
	"slices"

	"github.com/erikgeiser/promptkit/textinput"
)

const ResultTemplateTextInput = `
{{- print .Prompt " " (Foreground "206"  (Mask .FinalValue)) "\n" -}}
`

type Input struct {
	Arguments   [][]*string
	HiddenArgs  []*string
	Placeholder string
}

func (t Input) currentPlaceholder(name string) (string, error) {
	return " " + name + t.Placeholder, nil
}

func (t Input) Read() error {
	var input *textinput.TextInput
	var err error

	for _, arg := range t.Arguments {
		input = textinput.New(*arg[0])

		input.ResultTemplate = ResultTemplateTextInput

		hiddenInput := slices.Contains(t.HiddenArgs, arg[1])

		if hiddenInput {
			input.Hidden = true
		}

		input.Placeholder, _ = t.currentPlaceholder(*arg[0])
		*arg[1], err = input.RunPrompt()

		if err != nil {
			return err
		}
	}

	return nil
}
