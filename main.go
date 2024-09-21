package main

import (
	"errors"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
	"main.go/data/remote"
	"main.go/pkg/controllers/loader"
	"main.go/pkg/utils"
	"main.go/typingEngine"
	"os"
	"strconv"
)

var (
	selectedOption    string
	hasExitedPractice = false
)

func main() {
	database.InitializeDatabase()

	m := loader.InitialModel()
	p := tea.NewProgram(m)
	go func() {

		err := remote.FetchPractices()
		if err != nil {
			p.Send(loader.ErrMsg(err))
			return
		}

		p.Send(loader.DataLoadedMsg{})
		utils.ClearScreen()
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

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
				if selectedOption == "" {
					err := fmt.Sprintf("Please select an option to continue")
					return errors.New(err)
				}
				return nil
			}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if selectedOption == "Results" {
		if err = typingEngine.ReadMyResults(); err != nil {
			log.Fatal(err)
			return
		}
	} else {
		practice, err := strconv.ParseUint(selectedOption, 10, 32)

		if err = typingEngine.ReadPracticeLessons(uint(practice), &hasExitedPractice); err != nil {
			log.Fatal(err)
			return
		}
	}
}
