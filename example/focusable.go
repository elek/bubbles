package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elek/bubbles"
)

func NewPanel() tea.Model {
	h := ui.Horizontal{}
	m1 := &ui.Text{
		Content: func() string {
			return "first"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#0000EF")).BorderStyle(lipgloss.NormalBorder()),
	}
	m2 := &ui.Text{
		Content: func() string {
			return "first"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#EF00EF")).BorderStyle(lipgloss.NormalBorder()),
	}
	m3 := &ui.Text{
		Content: func() string {
			return "first"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#00EFEF")).BorderStyle(lipgloss.NormalBorder()),
	}

	v := ui.Vertical{}
	v.Add(m2, 10)
	v.Add(m3, 0)
	h.Add(m1, 100)
	h.Add(&v, 0)
	fg := ui.NewFocusGroup(&h)
	fg.Add(m1)
	fg.Add(m3)
	return fg
}

type StyledText struct {
}
