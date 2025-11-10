package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type DatabasePanel struct{}

func NewDatabasePanel() DatabasePanel { return DatabasePanel{} }

func (p DatabasePanel) Init() tea.Cmd { return nil }

func (p DatabasePanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p DatabasePanel) View() string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("[DatabasePanel] Placeholder content")
}
