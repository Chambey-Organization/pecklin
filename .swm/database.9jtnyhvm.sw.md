---
title: Database
---
<SwmSnippet path="/data/local/database/dao.go" line="10">

---

This function helps us mark a lesson as complete. If lesson result exists, update it with current stats like speed and if current speed is greater than the last best, set the best speed to the latest one. if lesson result is not in the database, create one.

```go

func CompleteLesson(progress *models.Progress) error {
	var existingProgress models.Progress
	DB.Where("lesson_id = ?", progress.LessonID).First(&existingProgress)

	if existingProgress.Id != 0 {
		existingProgress.CurrentSpeed = progress.CurrentSpeed

		if progress.CurrentSpeed > existingProgress.BestSpeed {
			existingProgress.BestSpeed = progress.CurrentSpeed
		}

		existingProgress.Complete = progress.Complete
		result := DB.Save(&existingProgress)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := DB.Create(progress)
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="36">

---

This function helps us retrieve the results for all lessons

```go
func GetResults() []models.Progress {
	var allProgress []models.Progress
	DB.Preload("Lesson").Find(&allProgress)
	return allProgress
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="42">

---

This function helps us to dd practices into the database. Mostly from the remote server.

```go
func InsertPractices(practices []models.Practice) {
	if len(practices) > 0 {
		DB.Delete(&models.Practice{})
		DB.Delete(&models.Lesson{})
		DB.Delete(&models.LessonContent{})
	}

	for _, practice := range practices {
		DB.Save(&practice)
	}
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="54">

---

This function helps us to select available practices in the database

```go
func ReadPractices() []models.Practice {
	var practices []models.Practice
	DB.Find(&practices)
	return practices
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="60">

---

This function selects all lesson with their respective

```go
func ReadPracticeLessons(practiceId uint) ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := DB.Preload("Content").Where("practice_id = ?", practiceId).Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="60">

---

&nbsp;contents.

```go
func ReadPracticeLessons(practiceId uint) ([]models.Lesson, error) {
	var lessons []models.Lesson
	if err := DB.Preload("Content").Where("practice_id = ?", practiceId).Find(&lessons).Error; err != nil {
		return nil, err
	}
	return lessons, nil
}
```

---

</SwmSnippet>

<SwmSnippet path="/data/local/database/dao.go" line="68">

---

This function helps us while debugging the code since we cannot directly print on the console, we write to a debug file which you can read. the file is called <SwmToken path="/data/local/database/dao.go" pos="69:13:14" line-data="	f, err := os.OpenFile(&quot;.debugFile&quot;, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)">`.debugFile`</SwmToken> on the root directory.

```go
func WriteToDebugFile(description string, input string) {
	f, err := os.OpenFile(".debugFile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf("%s %s \n", description, input))
}
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
