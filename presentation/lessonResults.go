package presentation

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"main.go/data/local/database"
	"os"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type tableModel struct {
	table table.Model
}

func (m tableModel) Init() tea.Cmd {
	return nil
}

func (m tableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			tea.Println("clicked on exit")
			navigation.Navigator.Pop()
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m tableModel) View() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33")).Render(" Results for your typing Lesson!")
	return "\n" + title + "\n\n" + baseStyle.Render(m.table.View()) + "\n  "
}

func LessonResultsPage(id uint) {

	result, err := database.GetLessonResultResult(id)

	if err != nil {
		tea.Println("Lesson not found")
	}

	columns := []table.Column{
		{Title: "Lesson", Width: 30},
		{Title: "Current Speed", Width: 30},
	}

	rows := []table.Row{
		{"Lesson", result.Lesson.Title},
		{"Current Speed", fmt.Sprintf("%.2f WPM", result.CurrentSpeed)},
		{"Best Speed", fmt.Sprintf("%.2f WPM", result.BestSpeed)},
		{"Accuracy", fmt.Sprintf("%.2f%%", result.Accuracy)},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(5),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	m := tableModel{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
