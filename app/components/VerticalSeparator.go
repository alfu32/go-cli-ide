package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type VerticalSeparator struct {
	Height int
	style  lipgloss.Style
}

func (a VerticalSeparator) Init() tea.Cmd { return nil }

func NewVerticalSeparator() VerticalSeparator {
	return VerticalSeparator{
		Height: 10,
		style:  lipgloss.NewStyle().Height(10).Background(lipgloss.Color("8")).Foreground(lipgloss.Color("8")),
	}
}

func (a VerticalSeparator) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	// switch _ := msg.(type) {
	// default:
	// case tea.KeyMsg:
	// 	// pass
	// }

	return a, cmd
}

func (a VerticalSeparator) View() string {
	sep := a.style.Height(a.Height).Render(" ")

	return sep
}
