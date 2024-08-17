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
	"main.go/pkg/utils/clear"
	"main.go/typingEngine"
	"os"
	"strconv"
)

var (
	practiceId string
)

func main() {
	m := loader.InitialModel()
	p := tea.NewProgram(m)

	go func() {
		database.InitializeDatabase()
		err := remote.FetchPractices()
		if err != nil {
			p.Send(loader.ErrMsg(err))
			return
		}

		p.Send(loader.DataLoadedMsg{})
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	practices := database.ReadPractices()

	var options []huh.Option[string]
	for _, practice := range practices {
		optionText := fmt.Sprintf("%d. %s", practice.ID, practice.Title)
		options = append(options, huh.NewOption(optionText, strconv.Itoa(int(practice.ID))))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Which typing practice do you want to practice today?").Options(
				options...,
			).Value(&practiceId).Validate(func(str string) error {
				if practiceId == "" {
					return errors.New("please select a lesson to continue")
				}
				return nil
			}),
		),
	)

	practice, err := strconv.ParseUint(practiceId, 10, 32)

	clear.ClearScreen()

	err = form.Run()

	if err != nil {
		log.Fatal(err)
	}

	err = typingEngine.ReadPracticeLessons(uint(practice))

	if err == nil {
		return
	} else {
		fmt.Printf("Exited lessons with error  %s\n", err.Error())
	}

}
