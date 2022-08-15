package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	ui "github.com/elek/bubbles"
	"os"
	"strings"
)

func main() {

	tabs := ui.NewTabs(
		ui.Tab{
			Name:  "filterable",
			Model: ui.NewFilterableList(createList("filtered ", 120)),
		},
		ui.Tab{
			Name:  "list",
			Model: ui.NewList(createList("simple ", 120)),
		},
	)
	if err := tea.NewProgram(ui.NewKillable(tabs)).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func createList(prefix string, size int) *StaticList {
	return &StaticList{
		prefix: prefix,
		size:   size,
	}

}

type StaticList struct {
	size   int
	prefix string
}

func (s *StaticList) Size() int {
	return s.size
}

func (s *StaticList) Render(i int) string {
	return fmt.Sprintf("%s #%d", s.prefix, i)
}

func (s *StaticList) Included(filter string, i int) bool {
	return strings.Contains(s.Render(i), filter)
}

func (s *StaticList) Item(i int) interface{} {
	return i
}

func (s *StaticList) Refresh() {

}
