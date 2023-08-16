package typing

import (
	"fmt"
	"github.com/eiannone/keyboard"
	typingSpeed2 "main.go/pkg/utils/typingSpeed"
	"os"
	"time"
)

const (
	introDelay = 1 * time.Second
)

func TypingPractice(sentences []string) {
	fmt.Println("Try this:")
	time.Sleep(introDelay)
	inputWords := ""

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	file, err := os.Open("sentences.txt")
	if err != nil {
		return
	}
	defer file.Close()

	startTime := time.Now()
	exitPractice := false // Variable to track if Esc key was pressed

	for _, sentence := range sentences {
		fmt.Printf("\n\n%s\n", sentence)
		var characters []rune
		var sentenceCharacters = []rune(sentence)

		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				break
			}

			if key == keyboard.KeyEnter {
				inputWords += " "
				break
			} else if key == keyboard.KeyEsc {
				fmt.Printf("\n\nExiting lesson 1 ...\n")
				exitPractice = true // Set the exit flag
				break
			} else if key == keyboard.KeySpace {
				inputWords += " "
				characters = append(characters, ' ')
			} else {
				inputWords += string(char)
				characters = append(characters, char)
			}

			if len(characters) > len(sentenceCharacters) {
				break
			}

			lastCharacter := characters[len(characters)-1]

			if characters[len(characters)-1] == sentenceCharacters[len(characters)-1] {
				fmt.Printf(string(lastCharacter))
			} else {
				fmt.Printf("^")
			}
		}

		if exitPractice { // Check if Esc key was pressed
			break // Exit the loop if Esc was pressed
		}
	}

	if !exitPractice {
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		typingSpeed := typingSpeed2.CalculateTypingSpeed(inputWords, duration)
		fmt.Printf("\n\nCongratulations! You have completed this lesson\nYour typing speed is: %.2f WPM\n", typingSpeed)
	}

}
