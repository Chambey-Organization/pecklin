package main

import (
	"bufio"
	"fmt"
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

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the path is a file and not a directory
		if !info.IsDir() {
			// Read the contents of the file into a string slice
			fileContent, err := readLinesFromFile(path)
			if err != nil {
				return err
			}

			// Extract the filename without extension from the path
			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			// Create a map with the filename as the title and the list of sentences as its value
			lessonData := models.Lesson{
				Title:   fileNameWithoutExt,
				Content: fileContent,
			}

			// Pass the map to the WelcomeScreen function
			welcome.WelcomeScreen(&lessonData)
			time.Sleep(3 * time.Second)
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
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
