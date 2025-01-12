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

type resultsTableModel struct {
	table table.Model
}

func (m resultsTableModel) Init() tea.Cmd {
	return nil
}

func (m resultsTableModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (m resultsTableModel) View() string {
	title := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33")).Render(" Results")
	return "\n" + title + "\n\n" + baseStyle.Render(m.table.View()) + "\n "
}

func ResultsPage() {

	results := database.GetResults()

	maxTitleLength := 0
	for _, result := range results {
		if len(result.Lesson.Title) > maxTitleLength {
			maxTitleLength = len(result.Lesson.Title)
		}
	}

	lessonColumnWidth := maxTitleLength + 2

	columns := []table.Column{
		{Title: "No.", Width: 5},
		{Title: "Lesson", Width: lessonColumnWidth},
		{Title: "Latest Speed", Width: 13},
		{Title: "BestSpeed", Width: 10},
		{Title: "Accuracy", Width: 10},
	}

	var rows []table.Row
	for index, result := range results {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d.", index+1),
			result.Lesson.Title,
			fmt.Sprintf("%.2f WPM", result.CurrentSpeed),
			fmt.Sprintf("%.2f WPM", result.BestSpeed),
			fmt.Sprintf("%.2f%%", result.Accuracy),
		})
	}

	tableHeight := len(rows) + 1

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithHeight(tableHeight), // Set the height dynamically
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

	m := resultsTableModel{t}
	if _, err := tea.NewProgram(m).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
