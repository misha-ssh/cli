package selection

import (
	"strconv"
	"strings"

	"github.com/erikgeiser/promptkit/selection"
	"github.com/muesli/termenv"
	"github.com/ssh-connection-manager/kernel/v2/pkg/output"
)

const (
	ResultChoiceName     = "alias"
	ResultTemplateSelect = `{{- print .Prompt " " (Foreground "206"  (alias .FinalChoice)) "\n" -}}`
)

type Selection struct {
	FilterPlaceholder string
	SelectionPrompt   string
	FilterPrompt      string
	Template          string
	PageSize          string
}

func (c Selection) Select(aliases []string) (string, error) {
	sp := selection.New(c.SelectionPrompt, aliases)

	sp.FilterPlaceholder = " " + c.FilterPlaceholder
	sp.FilterPrompt = c.FilterPrompt
	sp.Template = c.Template

	sp.ResultTemplate = ResultTemplateSelect

	pageSize, err := strconv.Atoi(c.PageSize)

	if err != nil {
		output.GetOutError("err page size must be an integer")
		return "", err
	}

	sp.PageSize = pageSize

	sp.Filter = func(filter string, choice *selection.Choice[string]) bool {
		return strings.HasPrefix(choice.Value, filter)
	}

	color := termenv.String().Foreground(termenv.ANSI256Color(206))

	sp.SelectedChoiceStyle = func(c *selection.Choice[string]) string {
		return color.Bold().Styled(c.Value)
	}

	sp.ExtendedTemplateFuncs = map[string]any{
		ResultChoiceName: func(c *selection.Choice[string]) string { return c.Value },
	}

	selectedValue, err := sp.RunPrompt()

	if err != nil {
		output.GetOutError("err running prompt")
		return "", err
	}

	return selectedValue, nil
}
