package main

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/huh"
	"github.com/eiannone/keyboard"
	"log"
	"main.go/database"
	"main.go/pkg/utils/clear"
	"main.go/typingEngine"
)

func main() {
	database.InitializeDatabase()

	var exitErr bool
	lessons := database.ReadCompletedLesson()
	allLessons := database.ReadAllLessons()

	var (
		lessonType string
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
	if exitErr {
		return
	}
	clear.ClearScreen()
	fmt.Println("\n Congratulations! You have completed all the lessons \n \nPress RETURN to redo the typing practice, SPACE to view lesson stats and ESC to quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyEnter {
			err := keyboard.Close()
			if err != nil {
				break
			}
			database.RedoLessons()
			lessons = database.ReadCompletedLesson()

			err = typingEngine.ReadTextLessons(lessons, &exitErr, lessonType)
			if exitErr {
				return
			}
			if err != nil {
				return
			}
		}

		if key == keyboard.KeySpace {
			for _, lesson := range allLessons {
				fmt.Printf("\nLesson Title: %s\n", lesson.Title)
				fmt.Printf("Typing Speed: %.2f WPM\n", lesson.BestSpeed)
				fmt.Println("---------------------------------")
			}
		}
		if key == keyboard.KeyEsc {
			break
		}
	}

	if err != nil {
		return
	}

}

/*


err := typingEngine.ReadTextLessons(lessons, &exitErr)
	if exitErr {
		return
	}
	clear.ClearScreen()
	fmt.Println("\n Congratulations! You have completed all the lessons \n \nPress RETURN to redo the typing practice, SPACE to view lesson stats and ESC to quit")

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyEnter {
			err := keyboard.Close()
			if err != nil {
				break
			}
			database.RedoLessons()
			lessons = database.ReadCompletedLesson()

			err = typingEngine.ReadTextLessons(lessons, &exitErr)
			if exitErr {
				return
			}
			if err != nil {
				return
			}
		}

		if key == keyboard.KeySpace {
			for _, lesson := range allLessons {
				fmt.Printf("\nLesson Title: %s\n", lesson.Title)
				fmt.Printf("Typing Speed: %.2f WPM\n", lesson.BestSpeed)
				fmt.Println("---------------------------------")
			}
		}
		if key == keyboard.KeyEsc {
			break
		}
	}

	if err != nil {
		return
	}
*/
