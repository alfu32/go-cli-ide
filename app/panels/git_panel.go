package panels

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type GitPanel struct{}

func NewGitPanel() GitPanel { return GitPanel{} }

func (p GitPanel) Init() tea.Cmd { return nil }

func (p GitPanel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return p, nil }

func (p GitPanel) View() string {
	return lipgloss.NewStyle().
		Padding(1, 2).
		Render("[GitPanel] Placeholder content")
}
