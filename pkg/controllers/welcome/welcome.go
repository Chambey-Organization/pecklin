package welcome

import (
	"fmt"

	"github.com/eiannone/keyboard"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/models"
	"main.go/pkg/utils/clear"
)

	func WelcomeScreen(lessonData *models.Lesson) {
	clear.ClearScreen()



    fmt.Printf("Welcome to lesson %s\n", lessonData.Title)
	fmt.Println("\nPress RETURN or SPACE to continue to typing practice. Press ESC to quit")


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

		if key == keyboard.KeySpace || key == keyboard.KeyEnter {
			err := keyboard.Close()
			if err != nil {
				break
			}
			clear.ClearScreen()
			typing.TypingPractice(lessonData)

		}

		if key == keyboard.KeyEsc {
			keyboard.Close()
			fmt.Println("Exiting lesson 1 ...")
			break
		}
	}

}


