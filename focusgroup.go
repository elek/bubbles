package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

func DefaultFocus(s lipgloss.Style) lipgloss.Style {
	w, h := s.GetFrameSize()
	st := s.BorderStyle(lipgloss.DoubleBorder())
	st = st.Height(s.GetHeight() - st.GetVerticalBorderSize() + w)
	st = st.Width(s.GetWidth() - st.GetHorizontalBorderSize() + h)
	return st
}

func DefaultDefocus(s lipgloss.Style) lipgloss.Style {
	w, h := s.GetFrameSize()
	st := s.BorderStyle(lipgloss.NormalBorder())
	st = st.Height(s.GetHeight() - st.GetVerticalBorderSize() + w)
	st = st.Width(s.GetWidth() - st.GetHorizontalBorderSize() + h)
	return st
}

type FocusMsg struct {
	Change  func(lipgloss.Style) lipgloss.Style
	Focused bool
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
	cmds := f.focusMessages()
	cmds = append(cmds, f.Model.Init())
	return tea.Batch(cmds...)
}

func (f *FocusGroup) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			f.Focused++
			if f.Focused > len(f.Items)-1 {
				f.Focused = 0
			}
			return f, tea.Batch(f.focusMessages()...)
		}
	}

	m, c := f.Model.Update(msg)
	f.Model = m
	return f, c
}

func (f *FocusGroup) focusMessages() []tea.Cmd {
	cmds := []tea.Cmd{}
	for i := 0; i < len(f.Items); i++ {
		changeMsg := FocusMsg{}
		changeMsg.Focused = i == f.Focused
		if i == f.Focused {
			changeMsg.Change = DefaultFocus
		} else {
			changeMsg.Change = DefaultDefocus
		}
		m, c := f.Items[i].Update(changeMsg)
		f.Items[i] = m
		cmds = append(cmds, c)
	}
	return cmds
}

func (f *FocusGroup) View() string {
	return f.Model.View()
}

func (f *FocusGroup) Add(m tea.Model) {
	f.Items = append(f.Items, m)
}

var _ tea.Model = &FocusGroup{}
