package presentation

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
	"main.go/pkg/utils"
)

func ResultsPage() {

	results := database.GetResults()

	utils.ClearScreen()
	var options []huh.Option[string]
	var number = 1
	for _, result := range results {
		optionText := fmt.Sprintf("%d. %s -> %.2f WPM (%.2f Accuracy) ", number, result.Lesson.Title, result.BestSpeed, result.Accuracy)
		options = append(options, huh.NewOption(optionText, result.Lesson.Title))
		number++
	}

	resultForm := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Results").Options(
				options...,
			),
		),
	)

	prog := tea.NewProgram(pageModel{form: resultForm})

	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
