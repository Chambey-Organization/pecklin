package presentation

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
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

	options = append(options, huh.NewOption("Exit", "Exit"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Results").Options(
				options...,
			),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if selectedOption == "Exit" {
		navigation.Navigator.Pop()
	}
}
