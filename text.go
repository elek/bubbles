package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Text struct {
	Content func() string
	size    Size
	Cached  string
	Style   lipgloss.Style
}

func NewText(content func() string) *Text {
	return &Text{
		Content: content,
		Style:   lipgloss.NewStyle(),
	}
}

var _ tea.Model = &Text{}

func (t *Text) Init() tea.Cmd {
	return func() tea.Msg {
		return RefreshMsg{}
	}
}

func (t *Text) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		t.size = NewSizeFromSizeMsg(msg)
		fx, fy := t.Style.GetFrameSize()
		t.Style = t.Style.Width(msg.Width - fx).Height(msg.Height - fy)

	case RefreshMsg:
		t.Refresh()
	}
	return t, nil
}

func (t *Text) View() string {
	out := ""
	fx, fy := t.Style.GetFrameSize()
	width := t.size.Width - fx
	height := t.size.Height - fy

	for ix, line := range strings.Split(t.Cached, "\n") {
		if ix > height-1 {
			break
		}
		if lipgloss.Width(line) > width {
			line = line[:width]
		}
		if out != "" {
			out += "\n"
		}
		out += line
	}
	return t.Style.Render(out)
}

func (t *Text) Refresh() {
	t.Cached = t.Content()
}
