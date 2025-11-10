package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type RequestsPanel struct{}

func NewRequestsPanel() RequestsPanel { return RequestsPanel{} }

func (p RequestsPanel) Init() tea.Cmd { return nil }

func (p RequestsPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p RequestsPanel) View() string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("[RequestsPanel] Placeholder content")
}
