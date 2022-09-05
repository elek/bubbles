package ui

import tea "github.com/charmbracelet/bubbletea"

func AsCommand(msg tea.Msg) tea.Cmd {
	return func() tea.Msg {
		return msg
	}
}
