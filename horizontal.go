package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
)

type sizeFunc func(available int) (size int)

type Horizontal struct {
	Size     tea.WindowSizeMsg
	Children Panels
}

type Panel struct {
	tea.Model
	Size SizeDefinition
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
		w := msg.Width
		h := msg.Height
		sizes := Resolve(v.Children.sizeDefinitions(), w)
		for i, p := range v.Children {
			nm, c := p.Update(tea.WindowSizeMsg{
				Width:  sizes[i],
				Height: h,
			})
			cmds = append(cmds, c)
			p.Model = nm
		}
		return v, tea.Batch(cmds...)
	}

	var cmds []tea.Cmd
	for _, p := range v.Children {
		nm, c := p.Update(msg)
		cmds = append(cmds, c)
		p.Model = nm
	}
	return v, tea.Batch(cmds...)
}

func (v *Horizontal) View() string {
	renders := make([][]string, len(v.Children))
	for ix, m := range v.Children {
		renders[ix] = strings.Split(m.View(), "\n")
	}

	out := ""
	for i := 0; i < v.Size.Height; i++ {
		if out != "" {
			out += "\n"
		}
		for p, _ := range v.Children {
			part := ""
			if len(renders[p]) > i {
				part = renders[p][i]
			}
			out += part
		}

	}
	return out
}

func (v Panels) sizeDefinitions() (res []SizeDefinition) {
	for _, p := range v {
		res = append(res, p.Size)
	}
	return
}
func exactSize(part string, i int) string {
	for c := len(part); c < i; c++ {
		part += " "
	}
	return AnsiTrim(part, i)
}

func (v *Horizontal) Add(model tea.Model, size SizeDefinition) {
	v.Children = append(v.Children, &Panel{
		Model: model,
		Size:  size,
	})
}

var _ tea.Model = &Horizontal{}
