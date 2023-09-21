package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/eiannone/keyboard"
	"main.go/database"
	"main.go/pkg/controllers/welcome"
	"main.go/pkg/models"
	"main.go/pkg/utils/clear"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	database.InitializeDatabase()
	var exitErr bool
	lessons := database.ReadCompletedLesson()

	err := ReadIncompleteLessons(lessons, &exitErr)
	if exitErr {
		return
	}
	clear.ClearScreen()
	fmt.Println("\n Congratulations! You have completed all the lessons \n \nPress RETURN or SPACE to redo the typing practice. Press ESC to quit")
	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			break
		}

		if key == keyboard.KeySpace || key == keyboard.KeyEnter {
			err := keyboard.Close()
			if err != nil {
				break
			}
			database.RedoLessons()
			lessons = database.ReadCompletedLesson()

			err = ReadIncompleteLessons(lessons, &exitErr)
			if exitErr {
				return
			}
			if err != nil {
				return
			}
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	if err != nil {
		return
	}
}

func ReadIncompleteLessons(lessons []models.Lesson, exitErr *bool) error {
	root := "lessons"

	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		var hasExitedLesson bool

		if !info.IsDir() {
			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			if lessonComplete(fileNameWithoutExt, lessons) {
				return nil
			}
			fileContent, err := readLinesFromFile(path)
			if err != nil {
				return err
			}

			lessonData := models.Lesson{
				Title:   fileNameWithoutExt,
				Content: fileContent,
			}

			welcome.WelcomeScreen(&lessonData, &hasExitedLesson)

			//check if user exited the lesson
			if hasExitedLesson {
				*exitErr = true
				return errors.New("user exited the lesson")
			}
			time.Sleep(3 * time.Second)
		}
		return nil
	})
}

// compare if lesson is complete
func lessonComplete(lessonTitle string, lessons []models.Lesson) bool {
	for _, lesson := range lessons {
		if lesson.Title == lessonTitle {
			return true
		}
	}
	return false
}

// readLinesFromFile reads the lines from a file and returns them as a string slice.
func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
