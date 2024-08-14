package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
	"main.go/data/remote"
	"main.go/pkg/utils/clear"
	"main.go/typingEngine"
	"strconv"
)

var (
	practiceId string
	exitErr    bool
)

func main() {
	database.InitializeDatabase()
	err := remote.FetchPractices()
	if err != nil {
		log.Fatal(err.Error())
		return
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

	err = typingEngine.ReadPracticeLessons(&exitErr, uint(practice))

	if err == nil {
		return
	} else if err.Error() == "user exited the practice" {
		clear.ClearScreen()

		err = form.Run()

		if err != nil {
			log.Fatal(err)
		}

		err = typingEngine.ReadPracticeLessons(&exitErr, uint(practice))
	} else {
		fmt.Printf("exit error is %s", err.Error())
	}

}
