package presentation

import (
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

type pageModel struct {
	form   *huh.Form
	action func()
}

func (m pageModel) Init() tea.Cmd {
	return nil
}

func (m pageModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter:
			if m.action != nil {
				m.action()
				return m, tea.Quit
			}
		case tea.KeyEsc:
			navigation.Navigator.Pop()
			return m, tea.Quit
		}

	}

	m.form.Update(msg)
	return m, nil
}

func (m pageModel) View() string {
	return m.form.View()
}
