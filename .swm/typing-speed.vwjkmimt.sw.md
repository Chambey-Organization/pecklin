---
title: Typing speed
---
<SwmSnippet path="/pkg/controllers/typing/typing.go" line="11">

---

This function takes the input words and start time, calculates time taken and determines user's words per minute

```go
func DisplayTypingSpeed(startTime time.Time, inputWords string, lesson *models.Lesson, accuracy float64) string {
	endTime := time.Now()
	duration := endTime.Sub(startTime)
	currentTypingSpeed := utils.CalculateTypingSpeed(inputWords, duration)

	progressLesson := models.Lesson{
		ID:         lesson.ID,
		PracticeID: lesson.PracticeID,
		Title:      lesson.Title,
		Active:     lesson.Active,
		Content:    lesson.Content,
	}

	progress := models.Progress{
		CurrentSpeed: currentTypingSpeed,
		BestSpeed:    currentTypingSpeed,
		Accuracy:     accuracy,
		Complete:     true,
		Lesson:       progressLesson,
		LessonID:     lesson.ID,
	}

	if err := database.CompleteLesson(&progress); err != nil {
		return fmt.Sprintf("Error saving your progress: %v", err)
	}

	return fmt.Sprintf("\n\n Congratulations! You have completed lesson %s\n Your typing speed is: %.2f WPM with an accuracy of %.2f%%\n", lesson.Title, currentTypingSpeed, accuracy)
}
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
