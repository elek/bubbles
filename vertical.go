package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Panels []*Panel

type Vertical struct {
	Size     tea.WindowSizeMsg
	Children Panels
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
		w := msg.Width
		h := msg.Height
		sizes := Resolve(v.Children.sizeDefinitions(), h)
		for i, p := range v.Children {
			nm, c := p.Update(tea.WindowSizeMsg{
				Width:  w,
				Height: sizes[i],
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

func (v *Vertical) Add(model tea.Model, size SizeDefinition) {
	v.Children = append(v.Children, &Panel{
		Model: model,
		Size:  size,
	})
}

var _ tea.Model = &Vertical{}
