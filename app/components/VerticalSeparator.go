package components

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type VerticalSeparator struct {
	Height    int
	style     StyleSheet
	HasFocus  bool
	IsHovered bool
}

func (a VerticalSeparator) Init() tea.Cmd { return nil }

func NewVerticalSeparator() VerticalSeparator {
	vs := VerticalSeparator{
		Height:    10,
		style:     NewStyleSheet(),
		HasFocus:  false,
		IsHovered: false,
	}
	vs.style.Visible = vs.style.Visible.Background(lipgloss.Color("8"))
	return vs
}
func (a VerticalSeparator) SetFocus(value bool) {
	a.HasFocus = value
}
func (a VerticalSeparator) SetHover(value bool) {
	a.IsHovered = value
}
func (a VerticalSeparator) GetFocus() bool {
	return a.HasFocus
}
func (a VerticalSeparator) GetHover() bool {
	return a.IsHovered
}
func (a VerticalSeparator) Update(msg tea.Msg) (VerticalSeparator, tea.Cmd) {
	var cmd tea.Cmd
	// switch _ := msg.(type) {
	// default:
	// case tea.KeyMsg:
	// 	// pass
	// }

	return a, cmd
}

func (a VerticalSeparator) View() string {
	style := a.style.Visible
	if a.HasFocus {
		style = a.style.Focused
	} else if a.IsHovered {
		style = a.style.Hovered
	}
	sep := style.Height(a.Height).BorderLeft(true).Render(" ")

	return sep
}
