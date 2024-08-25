package progressBar

import (
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

type ProgressModel struct {
	Progress progress.Model
}

func NewProgressModel() ProgressModel {
	return ProgressModel{
		Progress: progress.New(progress.WithDefaultGradient()),
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return nil
}

func (m ProgressModel) Update(msg tea.Msg) (ProgressModel, tea.Cmd) {
	switch msg := msg.(type) {
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd
	default:
		return m, nil
	}
}

func (m ProgressModel) View() string {
	return m.Progress.View()
}
