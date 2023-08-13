package main

import (
	"bufio"
	"fmt"
	"main.go/typing"
	"os"
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

	// Call TypingPractice with the sentences slice
	typing.TypingPractice(sentences)
}
