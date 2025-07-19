package input

import (
	"fmt"
	fl "github.com/misha-ssh/cli/internal/component/filepicker"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	ChatLimit           = 255
	ChatLimitPrivateKey = 1024
	Width               = 200
	HiddenChar          = '*'
)

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	cursorStyle  = focusedStyle
	noStyle      = lipgloss.NewStyle()

	focusedButton = focusedStyle.Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initModel() model {
	fileds := &Fields{}
	m := model{
		inputs: make([]textinput.Model, fileds.count()),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.Cursor.Style = cursorStyle
		t.CharLimit = ChatLimit
		t.Width = Width

		switch i {
		case fieldNumberAlias:
			t.Focus()
			t.Placeholder = fileds.getNameByNumber(fieldNumberAlias)
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			//t.Prompt = lipgloss.NewStyle().Width(2).Render("🏷")
		case fieldNumberLogin:
			t.Placeholder = fileds.getNameByNumber(fieldNumberLogin)
			//t.Prompt = lipgloss.NewStyle().Width(2).Render("👤")
		case fieldNumberPassword:
			t.Placeholder = fileds.getNameByNumber(fieldNumberPassword)
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = HiddenChar
			//t.Prompt = lipgloss.NewStyle().Width(2).Render("🔐")
		case fieldNumberPort:
			t.Placeholder = fileds.getNameByNumber(fieldNumberPort)
			t.Validate = portValidate
			//t.Prompt = lipgloss.NewStyle().Width(2).Render("🔌")
		case fieldNumberPrivateKey:
			t.Placeholder = fileds.getNameByNumber(fieldNumberPrivateKey) + " - " + "Press Ctrl+O to select the file"
			t.CharLimit = ChatLimitPrivateKey
			t.Validate = fileExistsValidate
			//t.Prompt = lipgloss.NewStyle().Width(2).Render("🗝️")1
		}

		m.inputs[i] = t
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) updateModel(msg tea.Msg) (model, tea.Cmd) {
	var cmds = make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyTab, tea.KeyShiftTab, tea.KeyEnter, tea.KeyUp, tea.KeyDown:
			cmds := make([]tea.Cmd, len(m.inputs))

			for i := range m.inputs {
				if m.inputs[i].Err != nil {
					return m, tea.EnableReportFocus
				}
			}

			s := msg.Type

			if s == tea.KeyEnter && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

			if s == tea.KeyUp || s == tea.KeyShiftTab {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = focusedStyle
					m.inputs[i].TextStyle = focusedStyle
					continue
				}

				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = noStyle
				m.inputs[i].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		case tea.KeyCtrlO:
			s := msg.Type

			if s == tea.KeyCtrlO && m.focusIndex == fieldNumberPrivateKey {
				filepickerFields := fl.Run()

				m.inputs[fieldNumberPrivateKey].SetValue(filepickerFields.PrivateKey)
				m.inputs[fieldNumberPrivateKey].Focus()
				m.inputs[fieldNumberPrivateKey].PromptStyle = focusedStyle
				m.inputs[fieldNumberPrivateKey].TextStyle = focusedStyle

				return m.updateModel(msg)
			}
		default:
			return m.updateModel(msg)
		}
	}

	return m.updateModel(msg)
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())

		if m.inputs[i].Err != nil {
			b.WriteString(fmt.Sprintf("\n%s", m.inputs[i].Err))
		}

		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	_, _ = fmt.Fprintf(&b, "\n\n%s", *button)

	return b.String()
}
