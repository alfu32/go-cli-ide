package app

import (
	"fmt"
	"ide-tui/app/components"
	"ide-tui/app/panels"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Application struct {
	tabs       VerticalTabs
	editor     CodeEditor
	separator  components.VerticalSeparator
	width      int
	height     int
	dragging   bool
	dragStartX int
	mouseX     int
	alert      *components.Alert
	x          int
	y          int
}

func New() Application {
	alert := components.NewAlert()
	return Application{
		tabs: NewVerticalTabs(
			[]string{
				"Prj",
				"Git",
				"DBa",
				"RQs",
			},
			map[string]tea.Model{
				"Prj": panels.NewProjectPanel(),
				"Git": panels.NewGitPanel(),
				"DBa": panels.NewDatabasePanel(),
				"RQs": panels.NewRequestsPanel(),
			}),
		editor:    NewCodeEditor(),
		separator: components.NewVerticalSeparator(),
		alert:     alert,
	}
}

func (a Application) Init() tea.Cmd { return nil }

func (a Application) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	_, _ = a.alert.Update(msg)
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		a.width, a.height = msg.Width, msg.Height
		a.tabs.height = msg.Height
		a.editor.height = msg.Height
		a.separator.Height = msg.Height
		a.editor.width = a.width - a.tabs.width

	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {
		case "ctrl+a":
			{
				// imperative alert trigger
				return a, a.alert.Show("Action executed!")
			}

		// These keys should exit the program.
		case "ctrl+c", "q":
			return a, tea.Quit
		}
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			// detect click on separator
			if msg.X == a.tabs.width-2 {
				a.dragging = true
				a.dragStartX = msg.X
			}
		}
		if msg.Action == tea.MouseActionRelease && msg.Button == tea.MouseButtonLeft {
			a.dragging = false
		}
		if a.dragging && msg.Action == tea.MouseActionMotion && msg.Button == tea.MouseButtonLeft {
			delta := msg.X - a.dragStartX
			a.tabs.width += delta
			if a.tabs.width < 10 {
				a.tabs.width = 10
			}
			if a.tabs.width > a.width-20 {
				a.tabs.width = a.width - 20
			}
			a.dragStartX = msg.X
		}
		a.x = msg.X
		a.y = msg.Y

	default:
		a.tabs, cmd = a.tabs.Update(msg)
		a.editor, _ = a.editor.Update(msg)
	}

	return a, cmd
}

func (a Application) View() string {
	tabsView := a.tabs.View()
	editorView := a.editor.View()
	sep := a.separator.View()
	bottom := lipgloss.NewStyle().
		Foreground(lipgloss.Color("7")).
		Background(lipgloss.Color("8")).
		Width(a.width).
		Render(fmt.Sprintf("x:%d y;%d dragging: %t startx:%d tabs.width:%d", a.x, a.y, a.dragging, a.dragStartX, a.tabs.width))

	return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Top, tabsView, sep, editorView), bottom)
}
