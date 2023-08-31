package welcome

import (
	"bufio"
	"fmt"
	"github.com/eiannone/keyboard"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/utils/clear"
	"os"
)

func WelcomeScreen() {
	clear.ClearScreen()
	fmt.Println("Welcome to lesson1")
	fmt.Println("\nPress RETURN or SPACE to continue to typing practice. Press ESC to quit")

	//open keyboard instance to start reading user input
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	//Read a file containing all the sentences  for lesson 1
	file, err := os.Open("sentences.txt")
	if err != nil {
		return
	}
	defer file.Close()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeySpace || key == keyboard.KeyEnter {
			//close keyboard before navigating to start typing
			err := keyboard.Close()
			if err != nil {
				break
			}
			startTypingPractice(file)
		}

		if key == keyboard.KeyEsc {
			keyboard.Close()
			fmt.Println("Exiting lesson 1 ...")
			break
		}
	}

}

func startTypingPractice(file *os.File) {
	//Read the file provided and break the sentences into a list for easy manipulation
	scanner := bufio.NewScanner(file)
	var sentences []string
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}
	clear.ClearScreen()
	typing.TypingPractice(sentences)
}
