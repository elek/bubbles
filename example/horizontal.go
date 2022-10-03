package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ui "github.com/elek/bubbles"
)

func NewHorizontalPane() tea.Model {
	h := ui.Horizontal{}
	h.Add(&ui.Text{
		Content: func() string {
			return "bbbbb"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#0000EF")).BorderStyle(lipgloss.NormalBorder()),
	}, ui.RemainingSize())
	h.Add(&ui.Text{
		Content: func() string {
			return "123\n123\n12312312312312312222222222222222222222222222222222222222222222222222222222222222222222222222222222222333"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#0000EF")).BorderStyle(lipgloss.NormalBorder()),
	}, ui.FixedSize(20))
	return &h
}
