package app

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type CodeEditor struct {
	width   int
	height  int
	content string
}

func NewCodeEditor() CodeEditor {
	return CodeEditor{content: "// CodeEditor placeholder"}
}

func (e CodeEditor) Init() tea.Cmd { return nil }

func (e CodeEditor) Update(msg tea.Msg) (CodeEditor, tea.Cmd) { return e, nil }

func (e CodeEditor) View() string {
	return lipgloss.NewStyle().
		Height(e.height).
		Width(e.width).
		// Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")).
		Padding(1, 2).
		Render(e.content)
}
