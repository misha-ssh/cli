package filepicker

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/charmbracelet/bubbles/filepicker"
)

func Run() *Fields {
	fp := filepicker.New()
	fp.CurrentDirectory, _ = os.UserHomeDir()

	m := model{
		filepicker: fp,
	}
	tm, _ := tea.NewProgram(&m).Run()
	mm := tm.(model)

	return &Fields{
		PrivateKey: mm.selectedFile,
	}
}
