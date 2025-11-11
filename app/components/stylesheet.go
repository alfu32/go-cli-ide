package components

import "github.com/charmbracelet/lipgloss"

type StyleSheet struct {
	Visible lipgloss.Style
	Hovered lipgloss.Style
	Focused lipgloss.Style
}

func NewStyleSheet() StyleSheet {
	return StyleSheet{
		Visible: lipgloss.NewStyle(),
		Hovered: lipgloss.NewStyle().Background(lipgloss.Color("6")).Foreground(lipgloss.Color("1")),
		Focused: lipgloss.NewStyle().Background(lipgloss.Color("7")).Foreground(lipgloss.Color("2")),
	}
}
