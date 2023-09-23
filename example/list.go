package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
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
	l.ChangeStyle(func(orig lipgloss.Style) lipgloss.Style {
		return orig.BorderStyle(lipgloss.NormalBorder())
	})
	return ui.NewHeader("NAME", l)
}
