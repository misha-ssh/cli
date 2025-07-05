package table

import "github.com/charmbracelet/lipgloss/table"

type Tabler interface {
	Display(value [][]string, headers []string) *table.Table
}
