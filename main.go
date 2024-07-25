package main

import (
	"fmt"
	"github.com/eiannone/keyboard"
	"main.go/database"
	"main.go/pkg/utils/clear"
	"main.go/typingEngine"
)

func main() {
	database.InitializeDatabase()
	var exitErr bool
	lessons := database.ReadCompletedLesson()
	allLessons := database.ReadAllLessons()

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
}
