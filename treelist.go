package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type TreeList[L, T any] struct {
	*List[T]
	down       func(T) (L, bool)
	items      func(L) []T
	breadcrumb []L
}

func NewTreeList[L, T any](content L, render func(T) string, items func(L) []T, down func(T) (L, bool)) *TreeList[L, T] {
	return &TreeList[L, T]{
		List:       NewList[T](items(content), render),
		down:       down,
		items:      items,
		breadcrumb: []L{content},
	}
}

func (t *TreeList[L, T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEscape, tea.KeyLeft:
			if len(t.breadcrumb) > 1 {
				t.breadcrumb = t.breadcrumb[0 : len(t.breadcrumb)-1]
				t.List.content = t.items(t.breadcrumb[len(t.breadcrumb)-1])
			}
			t.List.selected = 0
			t.List.verticalStart = 0
			return t, AsCommand(FocusedItemMsg[T]{Item: t.List.content[t.List.selected]})
		case tea.KeyEnter:
			sel, ok := t.down(t.List.content[t.selected])
			if ok {
				t.breadcrumb = append(t.breadcrumb, sel)
				t.List.content = t.items(sel)
				t.List.verticalStart = 0
				t.List.selected = 0
			}
			return t, AsCommand(FocusedItemMsg[T]{Item: t.List.content[t.List.selected]})
		}
	}
	m, c := t.List.Update(msg)
	t.List = m.(*List[T])
	return t, c
}

func (t *TreeList[L, T]) CurrentLevel() L {
	return t.breadcrumb[len(t.breadcrumb)-1]
}

func (t *TreeList[L, T]) Selected() T {
	return t.content[t.selected]
}

func (t *TreeList[L, T]) NavigateTo(f func(check T) bool) {
	t.verticalStart = 0
	t.selected = 0
	root := t.breadcrumb[0]
	t.breadcrumb = []L{}
	t.find(root, f)
	if len(t.breadcrumb) == 0 {
		t.breadcrumb = []L{root}
	}
	t.List.content = t.items(t.breadcrumb[len(t.breadcrumb)-1])

}

func (t *TreeList[L, T]) find(node L, f func(check T) bool) {
	// check children
	for ix, i := range t.items(node) {

		// is this it?
		if f(i) {
			t.selected = ix

			t.breadcrumb = []L{node}
			return
		}

		// can we dig deeper
		l, ok := t.down(i)
		if ok {
			t.find(l, f)
			// we found it
			if len(t.breadcrumb) > 0 {
				t.breadcrumb = append([]L{node}, t.breadcrumb...)
				return
			}
		}
	}

}
