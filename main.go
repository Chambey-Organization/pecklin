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
	practiceId        string
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

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Main menu").Options(
				options...,
			).Value(&practiceId).Validate(func(str string) error {
				if practiceId == "" {
					err := fmt.Sprintf("Please select a lesson to continue")
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

	practice, err := strconv.ParseUint(practiceId, 10, 32)

	err = typingEngine.ReadPracticeLessons(uint(practice), &hasExitedPractice)

	if err = typingEngine.ReadPracticeLessons(uint(practice), &hasExitedPractice); err != nil {
		log.Fatal(err)
		return
	}

}
