package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type Horizontal struct {
	Size     tea.WindowSizeMsg
	Children []tea.Model
	Limits   []int
}

func (v *Horizontal) Init() tea.Cmd {
	var cmds []tea.Cmd
	for ix := 0; ix < len(v.Children); ix++ {
		cmds = append(cmds, v.Children[ix].Init())
	}
	return tea.Batch(cmds...)
}

func (v *Horizontal) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		var cmds []tea.Cmd
		v.Size = msg

		used := 0
		for ix := 0; ix < len(v.Children); ix++ {
			w := v.Limits[ix]
			remaining := v.Size.Width - used
			if remaining <= 0 {
				break
			}
			if w > remaining || w == 0 {
				w = remaining
			}
			nm, c := v.Children[ix].Update(tea.WindowSizeMsg{
				Width:  w,
				Height: v.Size.Height,
			})
			cmds = append(cmds, c)
			v.Children[ix] = nm
			used += w
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

func (v *Horizontal) View() string {
	renders := make([][]string, len(v.Children))
	for ix, m := range v.Children {
		renders[ix] = strings.Split(m.View(), "\n")
	}

	out := ""
	for i := 0; i < v.Size.Height-1; i++ {
		for p, _ := range v.Children {
			part := ""
			if len(renders[p]) > i {
				part = renders[p][i]
			}
			out += part
		}
		out += "\n"
	}
	return out
}

func exactSize(part string, i int) string {
	for c := len(part); c < i; c++ {
		part += " "
	}
	return AnsiTrim(part, i)
}

func (v *Horizontal) Add(model tea.Model, height int) {
	v.Children = append(v.Children, model)
	v.Limits = append(v.Limits, height)
}

var _ tea.Model = &Horizontal{}
