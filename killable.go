package ui

import tea "github.com/charmbracelet/bubbletea"

type Killable struct {
	tea.Model
	enabledInput bool
}

func NewKillable(delegate tea.Model) tea.Model {
	return &Killable{
		Model:        delegate,
		enabledInput: true,
	}
}

func (k *Killable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case GrabInput:
		k.enabledInput = false
	case ReleaseInput:
		k.enabledInput = true
	case tea.KeyMsg:
		if k.enabledInput {
			if msg.String() == "esc" || msg.String() == "q" {
				return k, tea.Quit
			}
		}
	}
	m, c := k.Model.Update(msg)
	k.Model = m
	return k, c
}
