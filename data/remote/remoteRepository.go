package remote

import (
	"encoding/json"
	"fmt"
	"main.go/data/local/database"
	"main.go/domain/models"
	"os"
)

func FetchPractices() error {

	/* <p>For now we work with the local json folder</p>

	response, err := http.Get("http://13.244.41.201:8000/api/practice")
	if err != nil {
		return readPracticesFromFile()
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return readPracticesFromFile()
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return readPracticesFromFile()
	}

	var getPracticeDTO models.GetPracticesDTO
	if err := json.Unmarshal(responseData, &getPracticeDTO); err != nil {
		return readPracticesFromFile()
	}

	database.InsertPractices(getPracticeDTO.Practices)
	*/

	return readPracticesFromFile()
}

func readPracticesFromFile() error {
	fileData, err := os.ReadFile("practices.json")
	if err != nil {
		return fmt.Errorf("failed to read practices.json: %v", err)
	}

	var getPracticeDTO models.GetPracticesDTO
	if err := json.Unmarshal(fileData, &getPracticeDTO); err != nil {
		return fmt.Errorf("failed to unmarshal JSON from practices.json: %v", err)
	}

	database.InsertPractices(getPracticeDTO.Practices)
	fmt.Println("Successfully loaded practices from practices.json.")
	return nil
}
