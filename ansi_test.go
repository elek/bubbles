package ui

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"testing"
)

func TestTrim(t *testing.T) {
	//fmt.Println(AnsiTrim("abcdefghijkl", 5))
	styled := lipgloss.NewStyle().Foreground(lipgloss.Color("#EF0011")).Bold(true).Render("abcdefgh")
	res := AnsiTrim(styled, 5)
	fmt.Println(res)
}
