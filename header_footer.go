package ui

import tea "github.com/charmbracelet/bubbletea"

func NewHeader(header string, model tea.Model) *Vertical {
	v := Vertical{}
	v.Add(NewText(func() string {
		return header
	}), fixed{Value: 1})
	v.Add(model, remaining{})
	return &v
}

func NewFooter(model tea.Model, footer string) tea.Model {
	v := Vertical{}
	v.Add(model, remaining{})
	v.Add(NewText(func() string {
		return footer
	}), fixed{Value: 1})
	return &v
}
