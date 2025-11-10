package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ProjectPanel struct{}

func NewProjectPanel() ProjectPanel { return ProjectPanel{} }

func (p ProjectPanel) Init() tea.Cmd { return nil }

func (p ProjectPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p ProjectPanel) View() string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("[ProjectPanel] Placeholder content")
}
