package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Text struct {
	Content func() string
	Width   int
	Height  int
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
	case ResizeMsg:
		w, h := t.Style.GetFrameSize()
		t.Width = msg.Width - w
		t.Height = msg.Height - h
	case RefreshMsg:
		t.Cached = t.Content()
	}
	return t, nil
}

func (t *Text) View() string {
	out := ""

	for ix, line := range strings.Split(t.Cached, "\n") {
		if ix > t.Height-1 {
			break
		}
		if len(line) > t.Width {
			line = line[:t.Width]
		}
		if out != "" {
			out += "\n"
		}
		out += line
	}
	return out
}
