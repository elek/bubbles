package ui

import tea "github.com/charmbracelet/bubbletea"

type Size struct {
	Width  int
	Height int
}

func NewSizeFromSizeMsg(msg tea.WindowSizeMsg) Size {
	return Size{
		Width:  msg.Width,
		Height: msg.Height,
	}
}
