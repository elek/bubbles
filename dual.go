package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Dual struct {
	left       tea.Model
	right      tea.Model
	leftStyle  lipgloss.Style
	rightStyle lipgloss.Style
}

func NewDual(left tea.Model, right tea.Model) *Dual {
	return &Dual{
		left:  left,
		right: right,
		leftStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")),
		rightStyle: lipgloss.NewStyle().
			BorderStyle(lipgloss.NormalBorder()).
			BorderForeground(lipgloss.Color("63")),
	}
}

func (d *Dual) Init() tea.Cmd {
	return tea.Batch(
		d.left.Init(),
		d.right.Init(),
	)
}

func (d *Dual) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ResizeMsg:
		lw, lh := d.leftStyle.GetFrameSize()
		rw, rh := d.rightStyle.GetFrameSize()
		d.leftStyle = d.leftStyle.Height(msg.Height - lh)
		d.leftStyle = d.leftStyle.Width(msg.Width/2 - lw)
		d.rightStyle = d.rightStyle.Height(msg.Height - rh)
		d.rightStyle = d.rightStyle.Width(msg.Width/2 - rw)

		lu, lc := d.left.Update(ResizeMsg{
			Height: msg.Height - lh,
			Width:  msg.Width/2 - lw,
		})
		d.left = lu

		ru, rc := d.right.Update(ResizeMsg{
			Height: msg.Height - rh,
			Width:  msg.Width/2 - rw,
		})
		d.right = ru

		return d, tea.Batch(lc, rc)

	}

	lu, lc := d.left.Update(msg)
	d.left = lu

	ru, rc := d.right.Update(msg)
	d.right = ru

	return d, tea.Batch(lc, rc)
}

func (d *Dual) View() string {
	leftLines := strings.Split(d.leftStyle.Render(d.left.View()), "\n")
	rightLines := strings.Split(d.rightStyle.Render(d.right.View()), "\n")

	height := len(leftLines)
	rightHeight := len(rightLines)
	if rightHeight > height {
		height = rightHeight
	}
	s := ""
	for i := 0; i < height; i++ {
		line := ""
		if i < len(leftLines) {
			line += leftLines[i]
		}
		if i < len(rightLines) {
			line += rightLines[i]
		}
		if s != "" {
			s += "\n"
		}
		s += line
	}
	return s
}

var _ tea.Model = &Dual{}
