package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Vertical struct {
	Size     tea.WindowSizeMsg
	Children []tea.Model
	Limits   []int
	Styles   []lipgloss.Style
}

func (v *Vertical) Init() tea.Cmd {
	var cmds []tea.Cmd
	for ix := 0; ix < len(v.Children); ix++ {
		cmds = append(cmds, v.Children[ix].Init())
	}
	return tea.Batch(cmds...)
}

func (v *Vertical) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmds []tea.Cmd
		v.Size = msg

		used := 0
		for ix := 0; ix < len(v.Children); ix++ {
			h := v.Limits[ix]
			remaining := v.Size.Height - used
			if remaining <= 0 {
				break
			}
			if h > remaining || h == 0 {
				h = remaining
			}
			nm, c := v.Children[ix].Update(tea.WindowSizeMsg{
				Height: h,
				Width:  v.Size.Width,
			})
			cmds = append(cmds, c)
			v.Children[ix] = nm
			used += h
		}
		return v, tea.Batch(cmds...)
	}
	var cmds []tea.Cmd
	for ix := 0; ix < len(v.Children); ix++ {
		nm, c := v.Children[ix].Update(msg)
		cmds = append(cmds, c)
		v.Children[ix] = nm
	}
	return v, tea.Batch(cmds...)
}

func (v *Vertical) View() string {
	out := ""
	for _, m := range v.Children {
		if len(out) > 0 {
			out += "\n"
		}
		out += m.View()
	}
	return out
}

func (v *Vertical) Add(model tea.Model, height int) {
	v.Children = append(v.Children, model)
	v.Limits = append(v.Limits, height)
}

var _ tea.Model = &Vertical{}
