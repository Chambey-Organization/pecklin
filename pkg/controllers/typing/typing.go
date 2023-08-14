package typing

import (
	"bufio"
	"fmt"
	typingSpeed2 "main.go/pkg/utils/typingSpeed"
	"os"
	"strings"
	"time"
)

const (
	introDelay    = 1 * time.Second
	feedbackDelay = 1 * time.Second
	exitCommand   = "exit"
)

func TypingPractice(sentences []string) {
	fmt.Println("Try this: \n")
	time.Sleep(introDelay)

	for _, sentence := range sentences {
		fmt.Println(sentence)

		for {
			startTime := time.Now()

			reader := bufio.NewReader(os.Stdin)
			userInput, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			userInput = strings.TrimSpace(userInput)

			if userInput == exitCommand {
				fmt.Println("Exiting typing practice.")
				return
			}

			endTime := time.Now()
			duration := endTime.Sub(startTime)
			typingSpeed := typingSpeed2.CalculateTypingSpeed(sentence, duration)

			if userInput == sentence {
				fmt.Printf("\nCorrect! Well done! Typing Speed: %.2f WPM\n", typingSpeed)
				break // Exit the loop on correct input
			} else {
				fmt.Printf("\nIncorrect. Try again. Typing Speed: %.2f WPM\n", typingSpeed)
				fmt.Println("Type this sentence:\n ")
				fmt.Println(sentence)

			}

			time.Sleep(feedbackDelay)
		}
	}

	fmt.Println("\nCongratulations! You have completed the touch typing practice.")
}
