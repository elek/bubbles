package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	ui "github.com/elek/bubbles"
	"os"
)

func main() {

	var g []string
	for i := 0; i < 100; i++ {
		g = append(g, fmt.Sprintf("Item# %d", i))
	}

	v := ui.Vertical{}
	v.Add(&ui.Text{
		Content: func() string {
			return "aaa1\n2\n3\n4\n5\n6\n7\nx8\n9\n0\n1\n2\n3\n4\n5\n"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#EF00EF")).BorderStyle(lipgloss.NormalBorder()),
	}, ui.FixedSize(10))
	v.Add(&ui.Text{
		Content: func() string {
			return "bbbbb"
		},
		Style: lipgloss.NewStyle().BorderBackground(ui.White).Background(lipgloss.Color("#0000EF")).BorderStyle(lipgloss.NormalBorder()),
	}, ui.RemainingSize())

	tabs := ui.NewTabs(
		ui.Tab{
			Name:  "horizontal",
			Model: NewHorizontalPane(),
			Key:   "h",
		},
		ui.Tab{
			Name:  "focus",
			Model: NewFocusablePane(),
			Key:   "o",
		},
		ui.Tab{
			Name:  "v",
			Model: &v,
			Key:   "v",
		},
		ui.Tab{
			Name: "filterable",
			Model: ui.NewFilterableList(g, func(a string) string {
				return ">>" + a + "<<"
			}, nil),
			Key: "f",
		},
		ui.Tab{
			Name:  "tree",
			Model: createTree(),
			Key:   "t",
		},
		//ui.Tab{
		//	Name: "master-detail",
		//	Model: ui.NewSplit(createTree(), ui.NewDetail(func(t string) string {
		//		return "Detail: " + t
		//	})),
		//	Key: "m",
		//},
		ui.Tab{
			Name:  "list",
			Model: NewListPane(),
			Key:   "l",
		},
	)
	if err := tea.NewProgram(ui.NewKillable(tabs)).Start(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

func createTree() *ui.TreeList[string, string] {
	return ui.NewTreeList[string, string]("root",
		func(t string) string {
			return t
		},
		func(t string) (res []string) {
			for i := 0; i < 10; i++ {
				res = append(res, fmt.Sprintf("%s %d", t, i))
			}
			return res
		},
		func(t string) (string, bool) {
			return t, len(t) < 10
		},
	)
}
