package remote

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main.go/data/local/database"
	"main.go/domain/models"
	"net/http"
)

func FetchPractices() error {
	response, err := http.Get("https://mula-52f57-default-rtdb.firebaseio.com/pecking/-O4ByTT-VjlzLrLVagoJ.json")
	if err != nil {
		return fmt.Errorf("failed to get response: %v", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %v", response.StatusCode)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %v", err)
	}

	var practices []models.Practice
	if err := json.Unmarshal(responseData, &practices); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	database.InsertPractices(practices)
	database.ReadPractices()
	lessons := database.ReadAllLessons()
	lessonContent := database.ReadAllLessonsContent()
	fmt.Printf("Lessons size %d\n", len(lessons))
	fmt.Printf("Lesson content size %d\n", len(lessonContent))

	return nil
}
