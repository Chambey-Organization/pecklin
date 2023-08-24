package typing

import (
	"fmt"
	"github.com/eiannone/keyboard"
	typingSpeed2 "main.go/pkg/utils/typingSpeed"
	"time"
)

const (
	delay = 1 * time.Second
)

func TypingPractice(sentences []string) {
	fmt.Println("Try this:")
	time.Sleep(delay)

	inputWords := "" //variable to merge all words in the sentences to be used in calculating typing speed

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	startTime := time.Now()
	exitPractice := false // Variable to track if Esc key was pressed

	for _, sentence := range sentences {
		fmt.Printf("\n\n%s\n", sentence)

		var inputCharacters []rune

		var sentenceCharacters = []rune(sentence)
		for {
			char, key, err := keyboard.GetKey()
			if err != nil {
				break
			}

			if key == keyboard.KeyEnter {
				break
			} else if key == keyboard.KeyEsc {
				fmt.Printf("\n\nExiting lesson 1 ...\n")
				exitPractice = true
				break
			} else if key == keyboard.KeySpace {
				inputWords += " "
				inputCharacters = append(inputCharacters, ' ')
			} else {
				inputWords += string(char)
				inputCharacters = append(inputCharacters, char)
			}

			if len(inputCharacters) > len(sentenceCharacters) {
				break
			}

			lastCharacter := inputCharacters[len(inputCharacters)-1]

			if inputCharacters[len(inputCharacters)-1] == sentenceCharacters[len(inputCharacters)-1] {
				fmt.Printf(string(lastCharacter))
			} else {
				fmt.Printf("^")
			}
		}

		if exitPractice { // Check if Esc key was pressed
			break // Exit the loop if Esc was pressed
		}
	}

	// if user didn't exit calculate typing speed
	if !exitPractice {
		endTime := time.Now()
		duration := endTime.Sub(startTime)
		typingSpeed := typingSpeed2.CalculateTypingSpeed(inputWords, duration)
		fmt.Printf("\n\nCongratulations! You have completed this lesson\nYour typing speed is: %.2f WPM\n", typingSpeed)
	}

}
