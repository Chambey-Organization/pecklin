package typing

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"main.go/database"
	"main.go/pkg/models"
	"main.go/pkg/utils/typingSpeed"
)

const (
	delay = 1 * time.Second
)

// TypingPractice starts a typing practice session for a given lesson.
func TypingPractice(lessonData *models.Lesson) {
	fmt.Println("Try this:")
	time.Sleep(delay)

	inputWords := ""

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	startTime := time.Now()
	exitPractice := false

	for _, sentence := range lessonData.Content {
		fmt.Printf("\n\n%s\n", sentence)

		inputWords, exitPractice = handleTypingInput(sentence, inputWords)

		if exitPractice { // Check if Esc key was pressed
			break // Exit the loop if Esc was pressed
		}
	}

	// if user didn't exit calculate typing speed
	if !exitPractice {
		displayTypingSpeed(startTime, inputWords, lessonData.Title)
	}
}

// handleTypingInput handles user input for a given sentence and returns updated inputWords and exitPractice flag.
func handleTypingInput(sentence string, inputWords string) (string, bool) {
	var inputCharacters []rune
	sentenceCharacters := []rune(sentence)

	for {
		char, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeyEnter {
			break
		} else if key == keyboard.KeyEsc {
			fmt.Printf("\n\nExiting lesson  ...\n")
			return inputWords, true
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

		if lastCharacter == sentenceCharacters[len(inputCharacters)-1] {
			fmt.Print(string(lastCharacter))
		} else {
			fmt.Printf("^")
		}
	}

	return inputWords, false
}

func displayTypingSpeed(startTime time.Time, inputWords string, lessonTitle string) {

	endTime := time.Now()
	duration := endTime.Sub(startTime)
	typingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)
	fmt.Printf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lessonTitle, typingSpeed)
	var lesson models.LessonDTO
	lesson.Speed = fmt.Sprintf("%.2f WPM", typingSpeed)
	lesson.Title = lessonTitle
	database.CompleteLesson(lesson)
}
