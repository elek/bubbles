package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Bordered struct {
	tea.Model
	style lipgloss.Style
}

func WithBorder(m tea.Model) tea.Model {
	return &Bordered{
		Model: m,
		style: lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()).BorderForeground(White),
	}
}

func (b *Bordered) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		x, y := b.style.GetFrameSize()
		b.style.Width(msg.Width - x)
		b.style.Height(msg.Height - y)
		m, c := b.Model.Update(tea.WindowSizeMsg{
			Width:  msg.Width - x,
			Height: msg.Height - y,
		})
		b.Model = m
		return b, c
	}

	m, c := b.Model.Update(msg)
	b.Model = m
	return b, c
}

func (b *Bordered) View() string {
	return b.style.Render(b.Model.View())
}
