package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Selected = lipgloss.NewStyle().Background(lipgloss.Color("#404040"))
)

type List[T any] struct {
	content         []T
	size            tea.WindowSizeMsg
	verticalStart   int
	horizontalStart int
	selected        int
	style           lipgloss.Style
	Render          func(T) string
	SelectedStyle   lipgloss.Style
	Focused         bool
}

type FocusedItemMsg[T any] struct {
	Item T
}

func NewList[T any](content []T, render func(T) string) *List[T] {
	return &List[T]{
		content: content,
		size: tea.WindowSizeMsg{
			Width:  100,
			Height: 50,
		},
		Render:        render,
		SelectedStyle: Selected,
		style:         lipgloss.NewStyle(),
		Focused:       true,
	}
}

func (t *List[T]) Init() tea.Cmd {
	return func() tea.Msg {
		if len(t.content) > t.selected {
			return FocusedItemMsg[T]{
				Item: t.content[t.selected],
			}
		}
		return nil
	}
}

func (t *List[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FocusMsg:
		t.style = msg.Change(t.style)
		t.Focused = msg.Focused
	case tea.WindowSizeMsg:
		w, h := t.style.GetFrameSize()
		t.style.Height(msg.Height - h)
		t.style.Width(msg.Width - w)
		t.size = msg
	case tea.KeyMsg:
		if t.Focused {
			switch msg.Type {
			case tea.KeyCtrlRight:
				t.horizontalStart += 3
			case tea.KeyCtrlLeft:
				t.horizontalStart -= 3
				if t.horizontalStart < 0 {
					t.horizontalStart = 0
				}
			case tea.KeyDown, tea.KeyPgDown:
				if len(t.content) == 0 {
					return t, nil
				}
				_, h := t.style.GetFrameSize()
				height := t.size.Height - h
				if msg.Type == tea.KeyDown {
					t.selected++
				} else {
					t.selected += height
				}

				if t.selected > len(t.content)-1 {
					t.selected = len(t.content) - 1
				}
				if t.selected-t.verticalStart >= height {
					t.verticalStart += height
				}

				return t, func() tea.Msg {
					return FocusedItemMsg[T]{
						Item: t.content[t.selected],
					}
				}
			case tea.KeyUp, tea.KeyPgUp:
				if len(t.content) == 0 {
					return t, nil
				}
				_, h := t.style.GetFrameSize()
				height := t.size.Height - h

				if msg.Type == tea.KeyUp {
					t.selected--
				} else {
					t.selected -= height
				}

				if t.selected < 0 {
					t.selected = 0
				}
				if t.selected-t.verticalStart < 0 {
					t.verticalStart -= height
				}
				return t, func() tea.Msg {
					return FocusedItemMsg[T]{
						Item: t.content[t.selected],
					}
				}
			}
		}

	}
	return t, nil
}

func (t *List[T]) View() string {
	w, h := t.style.GetFrameSize()
	height := t.size.Height - h
	width := t.size.Width - w

	out := ""
	for i := t.verticalStart; i < t.verticalStart+height; i++ {
		if i >= len(t.content) {
			break
		}

		line := t.Render(t.content[i])
		if len(line) > t.horizontalStart {
			line = line[t.horizontalStart:]
		}
		if lipgloss.Width(line) > width {
			line = AnsiTrim(line, width)
		}
		if out != "" {
			out += "\n"
		}
		if i == t.selected {
			line = t.SelectedStyle.Render(line)
		}
		out += line
	}
	return t.style.Render(out)
}

func (t *List[T]) Reset() {
	t.selected = 0
	t.verticalStart = 0
}

func (t *List[T]) SetContent(n []T) {
	t.content = n
}

func (t *List[T]) Empty() bool {
	return len(t.content) == 0
}

func (t *List[T]) Selected() T {
	return t.content[t.selected]
}

func (t *List[T]) ChangeStyle(f func(orig lipgloss.Style) lipgloss.Style) {
	t.style = f(t.style)
}

func (t *List[T]) Select(ix int) {
	t.selected = ix
}
