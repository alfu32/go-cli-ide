package components

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Alert struct {
	Text     string
	Visible  bool
	duration time.Duration
	timer    *time.Timer
	hasFocus bool
}

type alertShowMsg struct {
	Text string
}

func NewAlert() *Alert {
	return &Alert{duration: 3 * time.Second}
}

func (a Alert) Focus(value bool) {
	a.hasFocus = value
}

// --- External call (imperative) ---
func (a *Alert) Show(text string) tea.Cmd {
	return func() tea.Msg { return alertShowMsg{Text: text} }
}

// --- Bubbletea model interface ---

func (a *Alert) Init() tea.Cmd { return nil }

func (a *Alert) Update(msg tea.Msg) (*Alert, tea.Cmd) {
	switch msg := msg.(type) {
	case alertShowMsg:
		a.Text = msg.Text
		a.Visible = true

		// auto-hide after duration
		return a, tea.Tick(a.duration, func(time.Time) tea.Msg {
			return alertHideMsg{}
		})

	case alertHideMsg:
		a.Visible = false
	}
	return a, nil
}

type alertHideMsg struct{}

func (a *Alert) View() string {
	if !a.Visible {
		return ""
	}
	style := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("11")).
		Background(lipgloss.Color("0")).
		Padding(1, 3).
		Align(lipgloss.Center)
	return style.Render(a.Text)
}
