package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	Selected = lipgloss.NewStyle().Background(lipgloss.Color("#404040"))
)

type List[T any] struct {
	content       []T
	size          tea.WindowSizeMsg
	start         int
	selected      int
	style         lipgloss.Style
	Render        func(T) string
	SelectedStyle lipgloss.Style
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
	case tea.WindowSizeMsg:
		w, h := t.style.GetFrameSize()
		t.style.Height(msg.Height - h)
		t.style.Width(msg.Width - w)
		t.size = msg
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyDown, tea.KeyPgDown:
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
			if t.selected-t.start >= height {
				t.start += height
			}
			return t, func() tea.Msg {
				return FocusedItemMsg[T]{
					Item: t.content[t.selected],
				}
			}
		case tea.KeyUp, tea.KeyPgUp:
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
			if t.selected-t.start < 0 {
				t.start -= height
			}
			return t, func() tea.Msg {
				return FocusedItemMsg[T]{
					Item: t.content[t.selected],
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
	for i := t.start; i < t.start+height; i++ {
		if i >= len(t.content) {
			break
		}

		line := t.Render(t.content[i])
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
	t.start = 0
}

func (t *List[T]) SetContent(n []T) {
	t.content = n
}

func (t *List[T]) Selected() T {
	return t.content[t.selected]
}
