package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/elek/bubbles"
)

func NewFocusablePane2() tea.Model {

	//m1 := &ui.Text{
	//	Content: func() string {
	//		return "first"
	//	},
	//	Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#EF00EF")).BorderStyle(lipgloss.NormalBorder()),
	//}
	//
	m1 := ui.NewList([]string{"asd", "basd"}, func(s string) string {
		return s
	})
	m1.ChangeStyle(func(orig lipgloss.Style) lipgloss.Style {
		return orig.BorderStyle(lipgloss.NormalBorder())
	})

	m2 := &ui.Scrolled{
		Content: func() string {
			return "first\n1\n2\n3\n4\n5\n6\n\n66666\n6\n6\n6\n6\n6\n"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#00EFEF")).BorderStyle(lipgloss.NormalBorder()),
	}

	v := ui.Vertical{}
	v.Add(m1, ui.FixedSize(5))
	v.Add(m2, ui.RemainingSize())
	fg := ui.NewFocusGroup(&v)
	fg.Add(m1)
	fg.Add(m2)
	return fg
}
