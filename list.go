package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Selected = lipgloss.NewStyle().Background(lipgloss.Color("#EFEFEF")).Foreground(lipgloss.Color("#DD0000"))
)

type List struct {
	content  ListContent
	width    int
	height   int
	start    int
	selected int
	style    lipgloss.Style
}

type SelectedMsg struct {
	Item interface{}
}
type ListContent interface {
	Size() int
	Render(int) string
	Item(int) interface{}
	Refresh()
}

type TreeContent interface {
	ListContent
	Choose(int)
	Up()
}

func NewList(content ListContent) *List {
	return &List{
		content: content,
		width:   40,
		height:  20,
	}
}

func (t *List) Init() tea.Cmd {
	return func() tea.Msg {
		return RefreshMsg{}
	}
}

func (t *List) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {
	case ResizeMsg:
		w, h := t.style.GetFrameSize()
		t.style.Height(msg.Height - h)
		t.style.Width(msg.Width - w)

		t.width = msg.Width - w
		t.height = msg.Height - h

	case RefreshMsg:
		t.content.Refresh()
	case tea.KeyMsg:
		if msg.String() == "q" {
			return t, tea.Quit
		}
		switch msg.Type {
		case tea.KeyEscape:
			if tree, ok := t.content.(TreeContent); ok {
				tree.Up()
				return t, func() tea.Msg {
					return SelectedMsg{
						Item: tree.Item(t.selected),
					}
				}
			}
		case tea.KeyDown, tea.KeyPgDown:
			if msg.Type == tea.KeyDown {
				t.selected++
			} else {
				t.selected += t.height
			}

			if t.selected > t.content.Size()-1 {
				t.selected = t.content.Size() - 1
			}
			if t.selected-t.start >= t.height {
				t.start += t.height
			}
			return t, func() tea.Msg {
				return SelectedMsg{
					Item: t.content.Item(t.selected),
				}
			}
		case tea.KeyUp, tea.KeyPgUp:
			if msg.Type == tea.KeyUp {
				t.selected--
			} else {
				t.selected -= t.height
			}

			if t.selected < 0 {
				t.selected = 0
			}
			if t.selected-t.start < 0 {
				t.start -= t.height
			}
			return t, func() tea.Msg {
				return SelectedMsg{
					Item: t.content.Item(t.selected),
				}
			}
		case tea.KeyEnter:
			if tree, ok := t.content.(TreeContent); ok {
				tree.Choose(t.selected)
			}
		}

	}
	return t, nil
}

func (t *List) View() string {
	out := ""
	for i := t.start; i < t.start+t.height; i++ {
		if i >= t.content.Size() {
			break
		}
		line := t.content.Render(i)
		if len(line) > t.width {
			line = line[:t.width]
		}
		if out != "" {
			out += "\n"
		}
		val := fmt.Sprintf("%d: %s", i, line)
		if i == t.selected {
			val = Selected.Render(val)
		}
		out += val
	}
	return t.style.Render(out)
}
