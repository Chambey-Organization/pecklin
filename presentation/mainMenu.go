package presentation

import (
	"errors"
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
	"strconv"
)

var selectedOption string

func MainMenu() {
	practices := database.ReadPractices()

	var options []huh.Option[string]
	var number = 1
	for _, practice := range practices {
		optionText := fmt.Sprintf("%d. %s", number, practice.Title)
		options = append(options, huh.NewOption(optionText, strconv.Itoa(int(practice.ID))))
		number++
	}

	//Add results option
	optionText := fmt.Sprintf("%d. Results", number)
	options = append(options, huh.NewOption(optionText, "Results"))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Main Menu").Options(
				options...,
			).Value(&selectedOption).Validate(func(str string) error {
				if selectedOption == "3" {
					err := fmt.Sprintf("Comming soon")
					return errors.New(err)
				}
				return nil
			}),
		),
	)

	prog := tea.NewProgram(pageModel{form: form, action: func() {
		if selectedOption != "3" {
			navigateAfterValidation()
		}
	}})

	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}

}

func navigateAfterValidation() {
	if selectedOption == "Results" {
		navigation.Navigator.Navigate(func() {
			ResultsPage()
		})
	} else {
		practice, err := strconv.ParseUint(selectedOption, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		navigation.Navigator.Navigate(func() {
			PracticeLessons(uint(practice))
		})
	}
}
