package main

import (
	"errors"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/database"
	"main.go/pkg/utils/clear"
	"main.go/typingEngine"
)

func main() {
	database.InitializeDatabase()

	lessons := database.ReadCompletedLesson()
	//allLessons := database.ReadAllLessons()

	var (
		lessonType string
		exitErr    bool
	)

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Which typing practice do you want to practice today?").Options(
				huh.NewOption("1. Basic keyboard lessons", "lessons/basics"),
				huh.NewOption("2. Golang basic lessons", "lessons/programming/go"),
				huh.NewOption("3. Golang advanced lessons", "lessons/programming/go"),
			).Value(&lessonType).Validate(func(str string) error {
				if lessonType == "" {
					return errors.New("please select a lesson to continue")
				}
				return nil
			}),
		),
	)

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
