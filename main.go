package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	file, err := os.Open("sentences.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var sentences []string
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	fmt.Println("Welcome to Touch Typing Practice!")

	fmt.Println("Type the sentences below as accurately and quickly as you can.")

	time.Sleep(2 * time.Second)

	for _, sentence := range sentences {
		fmt.Println("Sentence to type:")
		fmt.Println(sentence)

		startTime := time.Now()

		fmt.Print("Your input: ")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		endTime := time.Now()

		duration := endTime.Sub(startTime)

		words := strings.Fields(sentence)
		wordCount := len(words)
		seconds := duration.Seconds()
		typingSpeed := float64(wordCount) / (seconds / 60.0)

		if userInput == sentence {

			fmt.Printf("Correct! Well done! Typing Speed: %.2f WPM\n", typingSpeed)
			time.Sleep(2 * time.Second)

		} else {
			fmt.Printf("Incorrect. Try again. Typing Speed: %.2f WPM\n", typingSpeed)

		}
		time.Sleep(2 * time.Second)
	}
	fmt.Println("\nCongratulations! You have completed the touch typing practice.")
}
