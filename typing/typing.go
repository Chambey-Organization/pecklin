package typing

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	introDelay    = 2 * time.Second
	feedbackDelay = 2 * time.Second
	exitCommand   = "exit"
)

func readSentences(scanner *bufio.Scanner) []string {
	var sentences []string
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}
	return sentences
}

func calculateTypingSpeed(sentence string, duration time.Duration) float64 {
	words := strings.Fields(sentence)
	wordCount := len(words)
	seconds := duration.Seconds()
	return float64(wordCount) / (seconds / 60.0)
}

func TypingPractice(sentences []string) {
	fmt.Println("Welcome to Touch Typing Practice!")
	fmt.Println("Type the sentences below as accurately and quickly as you can.")
	time.Sleep(introDelay)

	for _, sentence := range sentences {
		fmt.Println("Type this sentence:\n ")
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
			typingSpeed := calculateTypingSpeed(sentence, duration)

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
