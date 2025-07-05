package table

import (
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

type Table struct {
	HeaderStyle  lipgloss.Color
	OddCellStyle lipgloss.Color
	EvenRowStyle lipgloss.Color
	BorderStyle  lipgloss.Color
}

const (
	WidthTable   = 70
	TopPadding   = 0
	RightPadding = 1
	CellWidth    = 1
)

func (t *Table) Display(value [][]string, headers []string) *table.Table {
	render := lipgloss.NewRenderer(os.Stdout)

	CellStyle := render.NewStyle().
		Width(CellWidth).
		Align(lipgloss.Center).
		Padding(TopPadding, RightPadding)

	OddRowStyle := CellStyle.Foreground(t.OddCellStyle)
	EvenRowStyle := CellStyle.Foreground(t.EvenRowStyle)
	HeaderColumnsStyle := CellStyle.Foreground(t.HeaderStyle)

	BorderStyle := lipgloss.NewStyle().Foreground(t.BorderStyle)

	viewTable := table.New().
		Border(lipgloss.NormalBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			if row == -1 {
				return HeaderColumnsStyle
			} else if row%2 == 0 {
				return OddRowStyle
			}

			return EvenRowStyle
		}).
		Headers(headers...).
		Rows(value...).
		Width(WidthTable)

	return viewTable
}
