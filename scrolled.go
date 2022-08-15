package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Scrolled struct {
	Content func() string
	Width   int
	Height  int
	Lines   []string
	Start   int
	Style   lipgloss.Style
}

func NewScrolled(content func() string) *Scrolled {
	return &Scrolled{
		Content: content,
		Style:   lipgloss.NewStyle(),
	}
}

var _ tea.Model = &Scrolled{}

type RefreshMsg struct {
}

func (t *Scrolled) Init() tea.Cmd {
	return func() tea.Msg {
		return RefreshMsg{}
	}
}

func (t *Scrolled) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ResizeMsg:
		w, h := t.Style.GetFrameSize()
		t.Style.Height(msg.Height - h)
		t.Style.Width(msg.Width - w)

		t.Width = msg.Width - w
		t.Height = msg.Height - h

	case RefreshMsg:
		t.Lines = strings.Split(t.Content(), "\n")
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown:
			t.Start++
		case tea.KeyUp:
			if t.Start > 0 {
				t.Start--
			}
		case tea.KeyPgUp:
			t.Start -= t.Height
			if t.Start <= 0 {
				t.Start = 0
			}
		case tea.KeyPgDown, tea.KeySpace:
			t.Start += t.Height
			if t.Start > len(t.Lines)-t.Height+2 {
				t.Start = len(t.Lines) - t.Height + 2
			}
			if t.Start < 0 {
				t.Start = 0
			}
		case tea.KeyHome:
			t.Start = 0
		case tea.KeyEnd:
			t.Start = len(t.Lines) - t.Height + 2
		default:
			fmt.Println(msg.Type)
			fmt.Println(msg.String())
		}

	}

	return t, nil
}

func (t *Scrolled) View() string {
	out := ""
	for i := t.Start; i < t.Start+t.Height; i++ {
		if i >= len(t.Lines) {
			break
		}
		line := t.Lines[i]
		if len(line) > t.Width {
			line = line[:t.Width]
		}
		if out != "" {
			out += "\n"
		}
		out += fmt.Sprintf("%s", line)
	}
	return t.Style.Render(out)
}

func (t *Scrolled) ResetPoistion() {
	t.Start = 0
}
