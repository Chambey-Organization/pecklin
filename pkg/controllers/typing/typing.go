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
	startTime := time.Now()
	inputWords := ""
	for _, sentence := range sentences {
		fmt.Printf("\n%s\n", sentence)

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
		inputWords = inputWords + " " + userInput

		time.Sleep(feedbackDelay)

	}

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	typingSpeed := typingSpeed2.CalculateTypingSpeed(inputWords, duration)
	fmt.Printf("\nCongratulations! You have completed this leason \nYour typing speed is: %.2f WPM\n", typingSpeed)
}
