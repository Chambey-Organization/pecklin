package typingEngine

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
	"main.go/pkg/utils"
)

func ReadMyResults() error {

	results := database.GetResults()

	utils.ClearScreen()
	var options []huh.Option[string]
	var number = 1
	for _, result := range results {
		optionText := fmt.Sprintf("%d. %s -> %.2f WPM (%.2f Accuracy) ", number, result.Lesson.Title, result.BestSpeed, result.Accuracy)
		options = append(options, huh.NewOption(optionText, result.Lesson.Title))
		number++
	}

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
		return err
	}

	return nil

}
