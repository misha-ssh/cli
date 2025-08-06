package output

import "github.com/charmbracelet/lipgloss"

var (
	successColor = lipgloss.Color("#04B575")
)

func Success(s string) {
	style := lipgloss.NewStyle().Foreground(successColor)
	println(style.Render(s))
}
