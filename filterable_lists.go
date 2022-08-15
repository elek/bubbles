package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type FilterableList struct {
	list            *List
	input           textinput.Model
	textBoxStyle    lipgloss.Style
	filter          string
	activeFilter    bool
	filteredContent *contentDelegate
}

type FilterableListContent interface {
	ListContent
	Included(filter string, i int) bool
}

func (f *FilterableList) Init() tea.Cmd {
	return f.list.Init()
}

func (f *FilterableList) Update(msg tea.Msg) (m tea.Model, c tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if f.activeFilter {
			if msg.String() == "esc" {
				f.activeFilter = false
				f.input.Blur()
				f.input.SetValue("")
				f.filter = ""
				f.filteredContent.Refilter("")
				return f, nil
			}
			if msg.String() == "enter" {
				f.activeFilter = false
				f.input.Blur()
				return f, nil
			}
			f.input, c = f.input.Update(msg)
			f.filteredContent.Refilter(f.input.Value())
			return f, c
		}
		switch msg.String() {
		case "/":
			f.activeFilter = true
			f.filter = ""
			nm1, c := f.list.Update(ResizeMsg{
				Width:  f.list.width,
				Height: f.list.height - 1,
			})
			f.list = nm1.(*List)
			f.input.Focus()
			return f, c

		}

	}

	m, c = f.list.Update(msg)
	f.list = m.(*List)
	return f, c
}

func (f *FilterableList) View() string {
	if f.activeFilter {
		return f.input.View() + "\n" + f.list.View()
	}
	return f.list.View()
}

func NewFilterableList(content FilterableListContent) *FilterableList {
	ti := textinput.New()
	filteredContent := &contentDelegate{
		content: content,
	}
	return &FilterableList{
		input:           ti,
		list:            NewList(filteredContent),
		textBoxStyle:    lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder()),
		filteredContent: filteredContent,
	}
}

type contentDelegate struct {
	content FilterableListContent
	filter  string
	reindex []int
}

func (c *contentDelegate) Size() int {
	return len(c.reindex)
}

func (c *contentDelegate) Render(i int) string {
	return c.content.Render(c.reindex[i])
}

func (c *contentDelegate) Item(i int) interface{} {
	return c.content.Item(c.reindex[i])
}

func (c *contentDelegate) Refresh() {
	c.content.Refresh()
	c.Refilter(c.filter)
}

func (c *contentDelegate) Refilter(filter string) {
	c.filter = filter
	c.reindex = []int{}
	for i := 0; i < c.content.Size(); i++ {
		if c.content.Included(filter, i) {
			c.reindex = append(c.reindex, i)
		}
	}
}
