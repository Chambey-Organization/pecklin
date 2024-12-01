---
title: Practice Lessons page
---
<SwmSnippet path="/presentation/lessonsPage.go" line="12">

---

This page displays the lessons available for this practice. We display them on a huh form and the user is able to select the lesson as an option.&nbsp;

```go
func PracticeLessons(practiceId uint) {
	practiceLessons, err := database.ReadPracticeLessons(practiceId)
	if err != nil {
		log.Fatal(err)
	}

	var selectedLessonIndex int
	var options []huh.Option[int]
	for i, lesson := range practiceLessons {
		optionText := fmt.Sprintf("%d. %s", i+1, lesson.Title)
		options = append(options, huh.NewOption(optionText, i))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Select Lesson").
				Options(options...).
				Value(&selectedLessonIndex),
		),
	)

	prog := tea.NewProgram(pageModel{form: form, action: func() {
		lesson := practiceLessons[selectedLessonIndex]
		navigation.Navigator.Navigate(func() {
			TypingPage(lesson)
		})
	}})

	if err := prog.Start(); err != nil {
		log.Fatal(err)
	}
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/lessonsPage.go" line="33">

---

Once a lesson is selected, we navigate the user to the next page where he is able to start typing. Note that running this in a new bubbletea program helps us capture user click esc key and pop the back stack

```go

	prog := tea.NewProgram(pageModel{form: form, action: func() {
		lesson := practiceLessons[selectedLessonIndex]
		navigation.Navigator.Navigate(func() {
			TypingPage(lesson)
		})
	}})
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
