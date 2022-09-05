package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func DefaultFocus(s lipgloss.Style) lipgloss.Style {
	return s.BorderStyle(lipgloss.DoubleBorder())
}

func DefaultDefocus(s lipgloss.Style) lipgloss.Style {
	return s.BorderStyle(lipgloss.NormalBorder())
}

type StyleChangeMsg struct {
	Change func(lipgloss.Style) lipgloss.Style
}

type FocusGroup struct {
	Items   []tea.Model
	Focused int
	tea.Model
}

func NewFocusGroup(child tea.Model) *FocusGroup {
	return &FocusGroup{
		Items: make([]tea.Model, 0),
		Model: child,
	}
}
func (f *FocusGroup) Init() tea.Cmd {
	m, c := f.Items[f.Focused].Update(StyleChangeMsg{
		Change: DefaultFocus,
	})
	f.Items[f.Focused] = m
	return tea.Batch(f.Model.Init(), c)
}

func (f *FocusGroup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			cmds := []tea.Cmd{}
			f.Focused++
			if f.Focused > len(f.Items)-1 {
				f.Focused = 0
			}
			for i := 0; i < len(f.Items); i++ {
				changeMsg := StyleChangeMsg{}
				if i == f.Focused {
					changeMsg.Change = DefaultFocus
				} else {
					changeMsg.Change = DefaultDefocus
				}
				m, c := f.Items[i].Update(changeMsg)
				f.Items[i] = m
				cmds = append(cmds, c)
			}
			return f, tea.Batch(cmds...)
		}
	}

	m, c := f.Model.Update(msg)
	f.Model = m
	return f, c
}

func (f *FocusGroup) View() string {
	return f.Model.View()
}

func (f *FocusGroup) Add(m tea.Model) {
	f.Items = append(f.Items, m)
}

var _ tea.Model = &FocusGroup{}
