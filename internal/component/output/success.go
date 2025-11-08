package output

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	green = lipgloss.Color("#04B575")
)

func output(colorText string) {
	fmt.Println(colorText)
}

func Success(str ...string) {
	output(lipgloss.NewStyle().Foreground(green).Render(str...))
}
