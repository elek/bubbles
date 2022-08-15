package ui

import tea "github.com/charmbracelet/bubbletea"

type Killable struct {
	delegate tea.Model
}

func NewKillable(delegate tea.Model) tea.Model {
	return &Killable{
		delegate: delegate,
	}
}

func (k *Killable) Init() tea.Cmd {
	return k.delegate.Init()
}

func (k *Killable) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "esc" {
			return k, tea.Quit
		}

	}
	return k.delegate.Update(msg)
}

func (k *Killable) View() string {
	return k.delegate.View()
}
