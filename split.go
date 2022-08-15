package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Split struct {
	dual     *Dual
	selected int
}

func NewSplit(left tea.Model, right tea.Model) *Split {
	dual := NewDual(left, right)
	s := &Split{
		dual: dual,
	}
	dual.leftStyle = dual.leftStyle.BorderStyle(lipgloss.DoubleBorder())
	dual.rightStyle = dual.rightStyle.BorderStyle(lipgloss.NormalBorder())
	return s
}

func (d *Split) Init() tea.Cmd {
	return d.dual.Init()
}

func (d *Split) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			if d.selected == 1 {
				d.dual.leftStyle = d.dual.leftStyle.BorderStyle(lipgloss.DoubleBorder())
				d.dual.rightStyle = d.dual.rightStyle.BorderStyle(lipgloss.NormalBorder())
				d.selected = 0
			} else {
				d.dual.leftStyle = d.dual.leftStyle.BorderStyle(lipgloss.NormalBorder())
				d.dual.rightStyle = d.dual.rightStyle.BorderStyle(lipgloss.DoubleBorder())
				d.selected = 1
			}
		default:
			if d.selected == 0 {
				d.dual.left, c = d.dual.left.Update(msg)
				return d, c
			} else {
				d.dual.right, c = d.dual.right.Update(msg)
				return d, c
			}
		}
	}

	m, c = d.dual.Update(msg)
	d.dual = m.(*Dual)
	return d, c

}

func (d *Split) View() string {
	return d.dual.View()
}

var _ tea.Model = &Dual{}
