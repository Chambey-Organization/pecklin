package main

import (
	"bufio"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"time"

	"main.go/database"
	"main.go/pkg/controllers/welcome"
	"main.go/pkg/models"
)

func main() {
	database.InitializeDatabase()
	root := "lessons"
	var exitErr error

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		hasExitedLesson := false

		// Check if the path is a file and not a directory
		if !info.IsDir() {
			// Extract the filename without extension from the path
			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			lessons := database.ReadCompletedLesson()

			// Check if the lesson title exists in the lessons slice
			if lessonComplete(fileNameWithoutExt, lessons) {
				return nil
			}
			// Read the contents of the file into a string slice
			fileContent, err := readLinesFromFile(path)
			if err != nil {
				return err
			}

			// Create a map with the filename as the title and the list of sentences as its value
			lessonData := models.Lesson{
				Title:   fileNameWithoutExt,
				Content: fileContent,
			}

			// Pass the map to the WelcomeScreen function
			welcome.WelcomeScreen(&lessonData, &hasExitedLesson)

			//check if user exited the lesson
			if hasExitedLesson {
				exitErr = errors.New("user exited the lesson")
				return exitErr
			}
			time.Sleep(3 * time.Second)
		}
		return nil
	})

	if exitErr != nil {
		return
	}

	if err != nil {
		return
	}
}

// compare if lesson exists
func lessonComplete(lessonTitle string, lessons []models.LessonDTO) bool {
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
