package app

import (
	fmt "fmt"
	"math"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type VerticalTabs struct {
	width      int
	height     int
	active     int
	order      []string
	panels     map[string]tea.Model
	_tab_width int
}

func NewVerticalTabs(order []string, panels map[string]tea.Model) VerticalTabs {
	return VerticalTabs{
		width:      28,
		height:     20,
		order:      order,  // []string{},                 // []string{"Project", "Git", "Database", "Requests"},
		panels:     panels, // make(map[string]tea.Model), // make(map[string]tea.Model),
		_tab_width: 10,
	}
}
func (m VerticalTabs) GetWidth() int {
	return m.width + m._tab_width
}
func (m VerticalTabs) SetPanels(order []string, panels map[string]tea.Model) VerticalTabs {
	var maxLabel = 0
	for _, name := range m.order {
		maxLabel = int(math.Max(float64(maxLabel), float64((len(name)))))
	}
	m._tab_width = maxLabel
	m.order = order
	m.panels = panels
	return m
}

func (m VerticalTabs) Init() tea.Cmd { return nil }

func (m VerticalTabs) Update(msg tea.Msg) (VerticalTabs, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.active = (m.active - 1 + len(m.order)) % len(m.order)
		case "down":
			m.active = (m.active + 1) % len(m.order)
		}
	}

	activeName := m.order[m.active]
	if activePanel, ok := m.panels[activeName]; ok {
		updated, c := activePanel.Update(msg)
		m.panels[activeName] = updated
		cmd = c
	}

	return m, cmd
}

func (m VerticalTabs) View() string {
	// Tabs list
	var body string
	for i, name := range m.order {
		prefix := "  "
		if i == m.active {
			prefix = "> "
		}
		body += fmt.Sprintf("%s%s\n", prefix, name)
	}

	list := lipgloss.NewStyle().
		MarginTop(2).
		Width(m._tab_width - 2).
		Height(m.height - 2).
		BorderForeground(lipgloss.Color("8")).
		Render(body)
	var panelView string
	if len(m.order) == 0 {
		panelView = "[No panel loaded]"
	} else {
		// Active panel content
		activeName := m.order[m.active]
		if p, ok := m.panels[activeName]; ok {
			panelView = p.View()
		} else {
			panelView = "[No panel loaded]"
		}
	}

	// Layout
	panelStyle := lipgloss.NewStyle().
		Height(m.height).
		Width(m.width - m._tab_width).
		// Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("8")) //.
	//Padding(1, 2)

	joined := lipgloss.JoinHorizontal(lipgloss.Top, list, panelStyle.Render(panelView))
	return joined
}
