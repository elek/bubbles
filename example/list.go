package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/elek/bubbles"
)

func NewListPane() tea.Model {
	var g []string
	for i := 0; i < 100; i++ {
		g = append(g, fmt.Sprintf("Item# %d", i))
	}

	l := ui.NewList(g, func(a string) string {
		return ">" + a + "<"
	})
	return ui.WithBorder(ui.NewHeader("NAME", l))
}
