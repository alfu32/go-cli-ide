package app

import (
	"fmt"
	"ide-tui/app/components"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tabItem struct{ title string }

func (i tabItem) Title() string       { return i.title }
func (i tabItem) Description() string { return "" }
func (i tabItem) FilterValue() string { return i.title }

type VerticalTabs struct {
	width      int
	height     int
	active     int
	intent     int
	tabNames   []string
	panels     map[string]tea.Model
	_tab_width int
	style      components.StyleSheet
}

func NewVerticalTabs(tabNames []string, panels map[string]tea.Model) VerticalTabs {
	vtabs := VerticalTabs{
		width:      40,
		height:     20,
		tabNames:   tabNames, // []string{},                 // []string{"Project", "Git", "Database", "Requests"},
		panels:     panels,   // make(map[string]tea.Model), // make(map[string]tea.Model),
		_tab_width: 5,
		intent:     -1,
	}
	vtabs.style = components.NewStyleSheet()
	return vtabs
}
func (m VerticalTabs) GetWidth() int {
	return m.width
}

func (m VerticalTabs) Init() tea.Cmd { return nil }

func (m VerticalTabs) Update(msg tea.Msg) (VerticalTabs, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.active = (m.active - 1 + len(m.tabNames)) % len(m.tabNames)
		case "down":
			m.active = (m.active + 1) % len(m.tabNames)
		}
	case tea.MouseMsg:
		if msg.Action == tea.MouseActionPress && msg.Button == tea.MouseButtonLeft {
			if msg.X < m._tab_width && (msg.Y) < len(m.tabNames)*3 {
				m.active = msg.Y / 3
			}
		}
		if msg.X < m._tab_width && (msg.Y) < len(m.tabNames)*3 {
			m.intent = msg.Y / 3 // len(m.tabNames)
		} else {
			m.intent = -1
		}
	}

	activeName := m.tabNames[m.active]
	if activePanel, ok := m.panels[activeName]; ok {
		updated, c := activePanel.Update(msg)
		m.panels[activeName] = updated
		cmd = c
	}

	return m, cmd
}

func (m VerticalTabs) View() string {
	// Tabs list
	// Build tab list with proper styles
	var lines []string
	for i, name := range m.tabNames {
		text := fmt.Sprintf("%4s\n%4s \n%4s", " ", name, " ") //"+      +\n  " + name + "   \n+      +"
		switch {
		case i == m.active:
			lines = append(lines, m.style.Focused.Render(text))
		case i == m.intent:
			lines = append(lines, m.style.Hovered.Render(text))
		default:
			lines = append(lines, m.style.Visible.Render(text))
		}
	}

	list := lipgloss.NewStyle().
		Width(m._tab_width).
		Height(m.height).
		Padding(1, 0, 0, 0).
		Render(strings.Join(lines, "\n"))
	var panelView string
	if len(m.tabNames) == 0 {
		panelView = "[No panel loaded]"
	} else {
		// Active panel content
		activeName := m.tabNames[m.active]
		if p, ok := m.panels[activeName]; ok {
			panelView = p.View()
		} else {
			panelView = "[No panel loaded]"
		}
	}

	// Layout

	panelStyle := lipgloss.NewStyle().
		Height(m.height).
		Width(m.width - m._tab_width - 1)

	// Join list and panel side by side
	sep := components.NewVerticalSeparator()
	sep.Height = m.height
	return lipgloss.JoinHorizontal(
		lipgloss.Top,
		list,
		sep.View(),
		panelStyle.Render(panelView),
	)
}
