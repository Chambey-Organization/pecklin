package typingEngine

import (
	"bufio"
	"errors"
	"main.go/pkg/controllers/welcome"
	"main.go/pkg/models"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ReadTextLessons(lessons []models.Lesson, exitErr *bool, lessonType string) error {

	return filepath.Walk(lessonType, func(path string, info os.FileInfo, err error) error {
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
			} else {
				time.Sleep(3 * time.Second)
			}
		}
		return nil
	})
}

func lessonComplete(lessonTitle string, lessons []models.Lesson) bool {
	for _, lesson := range lessons {
		if lesson.Title == lessonTitle {
			return true
		}
	}
	return false
}

func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}
