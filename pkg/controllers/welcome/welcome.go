package welcome

import (
	"bufio"
	"fmt"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/utils/clear"
	"os"
)

const (
	exitCommand = "exit"
)

func WelcomeScreen() {
	clear.CallClear()
	fmt.Println("Welcome to lesson1")
	fmt.Println("\nPress RETURN or SPACE to continue to typing practice.")

	file, err := os.Open("sentences.txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		key, err := reader.ReadByte()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		switch key {
		case '\r', '\n', ' ': // RETURN, NEWLINE, SPACE
			startTypingPractice(file)
			return
		default:
			fmt.Println("Press RETURN or SPACE to continue to typing practice.")
		}
	}
}

func startTypingPractice(file *os.File) {
	scanner := bufio.NewScanner(file)
	var sentences []string
	for scanner.Scan() {
		sentences = append(sentences, scanner.Text())
	}

	clear.CallClear()
	typing.TypingPractice(sentences)
}
