---
title: Remote Repository
---
<SwmSnippet path="/data/remote/remoteRepository.go" line="12">

---

This function is responsible for fetching the practices from the server and storing them in our local database.&nbsp;

```go
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
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
