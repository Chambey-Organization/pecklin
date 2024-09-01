package loader

import (
	"fmt"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type ErrMsg error

type DataLoadedMsg struct{}

type Model struct {
	Spinner  spinner.Model
	Quitting bool
	Err      error
	Loaded   bool
}

func InitialModel() Model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	return Model{Spinner: s}
}

func (m Model) Init() tea.Cmd {
	return m.Spinner.Tick
}

func (m Model) Exit() {
	m.Quitting = true
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			m.Quitting = true
			return m, tea.Quit
		default:
			return m, nil
		}

	case ErrMsg:
		m.Err = msg
		return m, nil

	case DataLoadedMsg:
		m.Loaded = true
		return m, tea.Quit

	default:
		if m.Loaded {
			return m, tea.Quit
		}
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	}
}

func (m Model) View() string {
	if m.Err != nil {
		return m.Err.Error()
	}
	str := fmt.Sprintf("\n\n %s Setting up your workspace, please wait...\n\n", m.Spinner.View())
	if m.Quitting {
		return str + "\n"
	}
	return str
}
