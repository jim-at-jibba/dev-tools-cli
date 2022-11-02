package tui

import "github.com/charmbracelet/lipgloss"

// 4 = blue
// 6 = magenta
// 7 = white
var ContainerStyle = lipgloss.NewStyle().
	Padding(1, 2, 1, 2).
	Margin(2).
	Border(lipgloss.NormalBorder(), true).
	BorderForeground(lipgloss.Color("4"))

var ValueStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("4")).
	PaddingLeft(1)

var LabelStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("7")).
	PaddingLeft(1)

var Spacer = lipgloss.NewStyle().Height(1)
