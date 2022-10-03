package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Tabs struct {
	children    []*Tab
	selected    int
	width       int
	height      int
	initialized []bool
}

type Tab struct {
	Name  string
	Model tea.Model
	Key   string
}

type ActivateTabMsg struct {
	Name string
	Msg  tea.Msg
}

func NewTabs(tabs ...Tab) *Tabs {
	t := &Tabs{}
	for ix, tab := range tabs {
		if tab.Key == "" {
			tab.Key = fmt.Sprintf("f%d", len(t.children)+1)
		}
		t.children = append(t.children, &tabs[ix])
	}
	t.initialized = make([]bool, len(t.children))
	return t
}

var _ tea.Model = &Tabs{}

func (t *Tabs) Init() tea.Cmd {
	return t.selectChild(t.selected)
}

func (t *Tabs) UpdateAll(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	for ix, tab := range t.children {
		m, c := tab.Model.Update(msg)
		t.children[ix].Model = m
		cmd = tea.Batch(cmd, c)
	}
	return t, cmd
}

func (t *Tabs) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case ActivateTabMsg:
		for ix, tab := range t.children {
			if tab.Key == msg.Name || tab.Name == msg.Name {
				cmd := t.selectChild(ix)
				if msg.Msg != nil {
					updated, _ := t.children[t.selected].Model.Update(msg.Msg)
					t.children[t.selected].Model = updated

				}
				return t, cmd
			}
		}
	case tea.WindowSizeMsg:
		t.width = msg.Width
		t.height = msg.Height

		var cmd tea.Cmd
		t.selectChild(t.selected)
		return t, cmd
	case tea.KeyMsg:
		for ix, tab := range t.children {
			if tab.Key == msg.String() {
				cmd := t.selectChild(ix)
				return t, cmd

			}
		}

	}
	updated, cmd := t.children[t.selected].Model.Update(msg)
	t.children[t.selected].Model = updated
	return t, cmd
}

func (t *Tabs) selectChild(index int) tea.Cmd {
	t.selected = index
	var initCmd, cmd tea.Cmd
	t.children[t.selected].Model, cmd = t.children[t.selected].Model.Update(tea.WindowSizeMsg{
		Width:  t.width,
		Height: t.height - 1,
	})

	if !t.initialized[index] {
		initCmd = t.children[t.selected].Model.Init()
		t.initialized[index] = true
	}
	return tea.Batch(initCmd, cmd)
}

func (t *Tabs) View() string {
	names := []string{}
	for _, a := range t.children {
		names = append(names, fmt.Sprintf("%s:%s", a.Key, a.Name))
	}
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#010101")).
		Background(lipgloss.Color("#EFEFEF")).
		Render(strings.Join(names, "    ")) + "\n" + t.children[t.selected].Model.View()
}

func (t *Tabs) Add(s string, details tea.Model, key string) *Tab {
	tab := &Tab{
		Name:  s,
		Model: details,
		Key:   key,
	}
	t.children = append(t.children, tab)
	t.initialized = append(t.initialized, false)
	return tab
}
