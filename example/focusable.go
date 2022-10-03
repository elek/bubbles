package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elek/bubbles"
)

func NewFocusablePane() tea.Model {
	h := ui.Horizontal{}
	m1 := &ui.Scrolled{
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
	m3 := &ui.Scrolled{
		Content: func() string {
			return "first\n1\n2\n3\n4\n5\n6\n\n66666\n6\n6\n6\n6\n6\n"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#00EFEF")).BorderStyle(lipgloss.NormalBorder()),
	}

	v := ui.Vertical{}
	v.Add(m2, ui.FixedSize(10))
	v.Add(m3, ui.RemainingSize())
	h.Add(m1, ui.FixedSize(100))
	h.Add(&v, ui.RemainingSize())
	fg := ui.NewFocusGroup(&h)
	fg.Add(m1)
	fg.Add(m3)
	return fg
}

type StyledText struct {
}
