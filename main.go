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
)

var (
	lessonType string
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
	lessons := database.ReadCompletedLesson()

	var options []huh.Option[string]
	for _, practice := range practices {
		optionText := fmt.Sprintf("%d. %s", practice.ID, practice.Title)
		optionValue := fmt.Sprintf("lessons/practice/%d", practice.ID)
		options = append(options, huh.NewOption(optionText, optionValue))
	}
	//allLessons := database.ReadAllLessons()

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Which typing practice do you want to practice today?").Options(
				options...,
			).Value(&lessonType).Validate(func(str string) error {
				if lessonType == "" {
					return errors.New("please select a lesson to continue")
				}
				return nil
			}),
		),
	)

	if !exitErr {
		clear.ClearScreen()
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		err = typingEngine.ReadTextLessons(lessons, &exitErr, lessonType)
		if err != nil {
			return
		}
	} else {
		clear.ClearScreen()
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}

		err = typingEngine.ReadTextLessons(lessons, &exitErr, lessonType)
		if err != nil {
			return
		}
	}
}
