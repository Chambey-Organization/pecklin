package progressBar

import (
	"fmt"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type TickMsg int

type ProgressModel struct {
	Progress   progress.Model
	Value      float64
	Seconds    int
	TimerCount *uint
}

func NewProgressModel(timerCount *uint) ProgressModel {
	var value float64
	var seconds int

	if timerCount != nil {
		value = 1.0
		seconds = int(*timerCount)
	} else {
		value = 0.0
		seconds = 0
	}

	return ProgressModel{
		Progress:   progress.New(progress.WithDefaultGradient()),
		Seconds:    seconds,
		Value:      value,
		TimerCount: timerCount,
	}
}

func (m ProgressModel) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(1) // Send a tick every second
	})
}

func (m ProgressModel) Update(msg tea.Msg) (ProgressModel, tea.Cmd) {

	switch msg := msg.(type) {
	case TickMsg:
		if m.Seconds > 0 && m.TimerCount != nil {
			m.Seconds--
			m.Value = float64(m.Seconds) / float64(*m.TimerCount)
		}
		if m.Seconds == 0 {
			return m, nil
		}
		return m, tick()

	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd

	default:
		return m, nil
	}
}

func (m ProgressModel) View() string {
	if m.TimerCount != nil {
		return fmt.Sprintf(
			" Time left: %d seconds\n %s", m.Seconds, m.Progress.ViewAs(m.Value),
		)
	}

	return ""
}
