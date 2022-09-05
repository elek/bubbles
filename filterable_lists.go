package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type FilterableList[T any] struct {
	*List[T]
	fullContent  []T
	filterFunc   func(string, T) bool
	input        textinput.Model
	textBoxStyle lipgloss.Style
	filter       string
	activeFilter bool
}

func (f *FilterableList[T]) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if f.activeFilter {
			if msg.String() == "esc" {
				f.activeFilter = false
				f.input.Blur()
				f.input.SetValue("")
				f.filter = ""
				f.refilter("")
				return f, AsCommand(ReleaseInput{})
			}
			if msg.String() == "enter" {
				f.activeFilter = false
				f.input.Blur()
				return f, AsCommand(ReleaseInput{})
			}
			f.input, c = f.input.Update(msg)
			f.refilter(f.input.Value())
			return f, c
		}
		switch msg.String() {
		case "/":
			f.activeFilter = true
			f.filter = ""
			nm1, c := f.List.Update(tea.WindowSizeMsg{
				Width:  f.List.size.Width,
				Height: f.List.size.Height - 1,
			})
			f.List = nm1.(*List[T])
			f.input.Focus()
			return f, tea.Batch(c, AsCommand(GrabInput{}))

		}

	}

	m, c = f.List.Update(msg)
	f.List = m.(*List[T])
	return f, c
}

func (f *FilterableList[T]) View() string {
	if f.activeFilter {
		return f.input.View() + "\n" + f.List.View()
	}
	return f.List.View()
}

func (f *FilterableList[T]) refilter(s string) {
	var filtered []T
	for _, v := range f.fullContent {
		if f.filterFunc(s, v) {
			filtered = append(filtered, v)
		}
	}
	f.List.content = filtered
}

func NewFilterableList[T any](content []T, render func(T) string, filter func(string, T) bool) *FilterableList[T] {
	if filter == nil {
		filter = func(s string, t T) bool {
			return strings.Contains(render(t), s)
		}
	}
	return &FilterableList[T]{
		fullContent: content,
		List:        NewList(content, render),
		filterFunc:  filter,
		input:       textinput.New(),
	}
}
