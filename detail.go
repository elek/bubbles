package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Detail[T any] struct {
	*Text
	focused T
}

func NewDetail[T any](render func(T) string) *Detail[T] {
	d := &Detail[T]{}
	d.Text = NewText(func() string {
		return render(d.focused)
	})
	return d

}
func (t *Detail[T]) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case FocusedItemMsg[T]:
		t.focused = msg.Item
	}

	m, c := t.Text.Update(msg)
	t.Text = m.(*Text)
	return t, c
}
