package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestScrolled(t *testing.T) {
	s := NewScrolled(func() string {
		return "abcd\nqwe\naaaaaaa\n123"
	})

	initialMsg := s.Init()()
	update, _ := s.Update(initialMsg)
	s = update.(*Scrolled)

	update, _ = s.Update(tea.WindowSizeMsg{
		Width:  4,
		Height: 3,
	})
	s = update.(*Scrolled)

	require.Equal(t, `abcd
qwe
aaaa`, s.View())

	//scroll down with one line
	update, _ = s.Update(tea.KeyMsg{
		Type: tea.KeyDown,
	})
	s = update.(*Scrolled)

	require.Equal(t, `qwe
aaaa
123`, s.View())

}
