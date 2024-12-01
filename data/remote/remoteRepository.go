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
	
	response, err := http.Get("https://charlesmuchogo.com/api/practice")
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

	var getPracticeDTO models.GetPracticesDTO
	if err := json.Unmarshal(responseData, &getPracticeDTO); err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	database.InsertPractices(getPracticeDTO.Practices)
	return nil
}
