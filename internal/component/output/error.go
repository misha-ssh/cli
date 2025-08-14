package output

import "github.com/charmbracelet/lipgloss"

var (
	ErrorColor = lipgloss.Color("#FF0000")
)

// Error get message line for erros
func Error(s string) {
	style := lipgloss.NewStyle().Foreground(ErrorColor)
	println(style.Render(s))
}
