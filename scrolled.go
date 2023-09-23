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
	Focused bool
}

func NewScrolled(content func() string) *Scrolled {
	return &Scrolled{
		Content: content,
		Style:   lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()),
		Focused: true,
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
	case tea.WindowSizeMsg:
		w, h := t.Style.GetFrameSize()
		t.Style = t.Style.Height(msg.Height - h)
		t.Style = t.Style.Width(msg.Width - w)

		t.Width = msg.Width - w
		t.Height = msg.Height - h
	case FocusMsg:
		//t.Style = msg.Change(t.Style)
		t.Focused = msg.Focused
	case RefreshMsg:
		t.Lines = strings.Split(t.Content(), "\n")
	case tea.KeyMsg:
		if t.Focused {
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
			}
		}

	}

	return t, nil
}

func (t *Scrolled) ScrollToEnd() {
	for ; t.Start < len(t.Lines)-t.Height; t.Start += t.Height {
	}
	if t.Start > len(t.Lines)-t.Height+2 {
		t.Start = len(t.Lines) - t.Height + 2
	}
	if t.Start < 0 {
		t.Start = 0
	}
}

func (t *Scrolled) View() string {
	out := ""
	for i := t.Start; i < t.Start+t.Height; i++ {
		if i >= len(t.Lines) {
			break
		}
		line := t.Lines[i]
		if len(line) > t.Width {
			line = AnsiTrim(line, t.Width)
		}
		if out != "" {
			out += "\n"
		}
		out += fmt.Sprintf("%s", line)
	}
	return t.Style.Render(out)
}

func (t *Scrolled) ResetPosition() {
	t.Start = 0
}
