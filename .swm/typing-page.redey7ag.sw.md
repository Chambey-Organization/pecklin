---
title: 'Typing Page  '
---
<SwmSnippet path="/presentation/typingPage.go" line="30">

---

We initialize this model to help us manage the various moving parts of our touch typing page like text styles, height and width of typing area, progress bar lesson, prompts, timer and errors.

```go
type model struct {
	viewport         viewport.Model
	prompt           []string
	input            string
	textAreaValue    string
	textarea         textarea.Model
	titleStyle       lipgloss.Style
	promptStyle      lipgloss.Style
	resultsStyle     lipgloss.Style
	err              error
	lesson           *models.Lesson
	prompts          []models.LessonContent
	currentIndex     int
	startTime        time.Time
	totalAccuracy    float64
	hasStartedTyping bool
	progress         progressBar.ProgressModel
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="49">

---

&nbsp;

```go
func TypingPage(lesson models.Lesson) {
	p := tea.NewProgram(initialModel(lesson))

	if finalModel, err := p.Run(); err != nil {
		log.Fatal(err)
	} else if lessonModel, ok := finalModel.(model); ok && lessonModel.err == nil {
		time.Sleep(time.Second * 7)
		navigation.Navigator.Pop()
	}
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="49">

---

This is the entry point of our app. It helps us manage navigation actions on our page. if lesson is complete or theres an error on this psge (User pressing "esc" is returned as an error), navigate back.

```go
func TypingPage(lesson models.Lesson) {
	p := tea.NewProgram(initialModel(lesson))

	if finalModel, err := p.Run(); err != nil {
		log.Fatal(err)
	} else if lessonModel, ok := finalModel.(model); ok && lessonModel.err == nil {
		time.Sleep(time.Second * 7)
		navigation.Navigator.Pop()
	}
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="60">

---

We initialize the model struct here giving it the default values.

```go
func initialModel(lesson models.Lesson) model {
	ta := textarea.New()
	ta.Placeholder = "Type the prompt"
	ta.Focus()

	ta.Prompt = " > "

	ta.SetWidth(100)
	ta.SetHeight(1)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#211efb"))
	var input []string

	titleText := fmt.Sprintf(" Welcome to lesson %s", lesson.Title)
	input = append(input)
	title := fmt.Sprintf(titleStyle.Render(titleText))

	vp := viewport.New(100, 17)

	vp.SetContent(titleStyle.Render(title))

	ta.KeyMap.InsertNewline.SetEnabled(false)

	prog := progress.New(progress.WithDefaultGradient())
	prog.Width = 50

	var progressBarModel progressBar.ProgressModel

	if lesson.TimerCount != nil {
		progressBarModel = progressBar.NewProgressModel(lesson.TimerCount)
	}

	m := model{
		textarea:         ta,
		prompt:           []string{},
		viewport:         vp,
		err:              nil,
		input:            "",
		textAreaValue:    "",
		totalAccuracy:    0,
		lesson:           &lesson,
		prompts:          lesson.Content,
		titleStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color("#211efb")).Bold(true),
		promptStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("#53C2C5")).Bold(true),
		resultsStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("#f817b0")).Bold(true),
		currentIndex:     0,
		hasStartedTyping: false,
		progress:         progressBarModel,
	}

	if len(m.lesson.Content) > 0 {
		m.progress.Progress.Width = 50
		prompt := m.lesson.Content[m.currentIndex].Prompt
		m.input = prompt
		m.prompt = append(m.prompt, " Prompt: "+m.promptStyle.Render(prompt))
		m.viewport.SetContent(strings.Join(m.prompt, "\n"))
		m.startTime = time.Now()

	}

	return m
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="114">

---

If the lesson content is not empty we set the first prompt here, an record the start time.&nbsp;

```go
	if len(m.lesson.Content) > 0 {
		m.progress.Progress.Width = 50
		prompt := m.lesson.Content[m.currentIndex].Prompt
		m.input = prompt
		m.prompt = append(m.prompt, " Prompt: "+m.promptStyle.Render(prompt))
		m.viewport.SetContent(strings.Join(m.prompt, "\n"))
		m.startTime = time.Now()

	}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="127">

---

This is the initial model of bubble tea to initialize the typing page. we set the text area to blink as a call to action to type the prompt

```go
func (m model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		m.progress.Init(),
	)
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="134">

---

This is an inbuilt callback function by bubble tea to get the values and process them while the user typing. Let's break down how we process the values down below.

```go
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd   tea.Cmd
		vpCmd   tea.Cmd
		progCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case progressBar.TickMsg:

		m.progress, progCmd = m.progress.Update(msg)
		prog := fmt.Sprintf("%f value and seconds is %d", m.progress.Value, m.progress.Seconds)
		database.WriteToDebugFile("TickMsg with values ->", prog)
		return m, tea.Batch(tiCmd, vpCmd, progCmd)
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			navigation.Navigator.Pop()
			m.err = errors.New("exited lesson")
			return m, tea.Quit
		case tea.KeyBackspace:
			if len(m.textAreaValue) > 0 && m.currentIndex < len(m.prompts) {
				m.textAreaValue = m.textAreaValue[:len(m.textAreaValue)-1]
				prompt := m.prompts[m.currentIndex].Prompt

				// Compare the input and get the highlighted input
				highlightedInput, _ := CompareAndHighlightInput(m.textAreaValue, prompt)
				m.textarea.Reset()
				m.textarea.Prompt = highlightedInput

				typingProgress := float64(len(m.input)) / float64(len(prompt))
				m.progress.Progress.SetPercent(typingProgress)
			}
		case tea.KeyEnter:
			input := m.textAreaValue

			if len(input) > 0 {
				prompt := m.prompts[m.currentIndex].Prompt
				highlightedInput, accuracy = CompareAndHighlightInput(input, prompt)
				m.prompt = append(m.prompt, fmt.Sprintf(" Input: %s (%.2f%% correct)\n", highlightedInput, accuracy))
				m.totalAccuracy += accuracy
			}

			m.textarea.Reset()
			m.textarea.Blur()
			m.viewport.GotoTop()
			m.textarea.Focus()
			m.textarea.Prompt = "> "

			// reset the text area value to blank
			m.textAreaValue = ""
			//Reintroduce the placeholder
			m.textarea.Placeholder = "Type the prompt"

			m.currentIndex++
			if m.currentIndex < len(m.prompts) {

				prompt := m.prompts[m.currentIndex].Prompt
				typingProgress := float64(len(input)) / float64(len(prompt))
				m.progress.Progress.SetPercent(typingProgress)

				m.prompt = append(m.prompt, " Prompt: "+m.promptStyle.Render(prompt))

				m.input = fmt.Sprintf("%s %s", m.input, prompt)
				m.viewport.SetContent(strings.Join(m.prompt, "\n"))

			} else {
				averageAccuracy := m.totalAccuracy / float64(len(m.prompts))
				database.WriteToDebugFile("m.input is while displaying ->", m.input)
				m.textarea.Prompt = " "
				m.textarea.Placeholder = ""
				m.prompt = append(m.prompt, m.resultsStyle.Render(typing.DisplayTypingSpeed(m.startTime, m.input, m.lesson, averageAccuracy)))
				m.viewport.SetContent(strings.Join(m.prompt, "\n"))
				return m, tea.Quit
			}
		}

	default:
		// Validate while typing for real-time feedback
		input := m.textarea.Value()

		m.textAreaValue = m.textAreaValue + input

		if len(input) > 0 && m.currentIndex < len(m.prompts) {
			prompt := m.prompts[m.currentIndex].Prompt

			// Compare the input and get the highlighted input
			highlightedInput, _ := CompareAndHighlightInput(m.textAreaValue, prompt)
			m.textarea.Reset()
			m.textarea.Placeholder = ""
			m.textarea.Prompt = highlightedInput

			typingProgress := float64(len(m.input)) / float64(len(prompt))
			m.progress.Progress.SetPercent(typingProgress)
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="135">

---

We initialize these variables and set the text area values to the received text. We also update the viewport to show the update.

```go
	var (
		tiCmd   tea.Cmd
		vpCmd   tea.Cmd
		progCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="143">

---

We now have the event/message update and we can process is accordingly. We use this switch statement to differentiate between ticking message, key message and error. in case of progress bar message, we update the progress bar&nbsp;

```go

	switch msg := msg.(type) {
	case progressBar.TickMsg:

		m.progress, progCmd = m.progress.Update(msg)
		prog := fmt.Sprintf("%f value and seconds is %d", m.progress.Value, m.progress.Seconds)
		database.WriteToDebugFile("TickMsg with values ->", prog)
		return m, tea.Batch(tiCmd, vpCmd, progCmd)
	case tea.KeyMsg:
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="151">

---

In case of key message, we can get the specific key pressed and e process accordingly.

```go
	case tea.KeyMsg:
		switch msg.Type {
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="152">

---

In case of esc button we throw an error and navigate the user from this page. The error helps us stop further execution of this page and manage resources efficiently.

```go
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			navigation.Navigator.Pop()
			m.err = errors.New("exited lesson")
			return m, tea.Quit
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="157">

---

in case of a backspace, if the user has typed something on the text area, we remove the last character and reformat the typed text. then update the text area.

```go
		case tea.KeyBackspace:
			if len(m.textAreaValue) > 0 && m.currentIndex < len(m.prompts) {
				m.textAreaValue = m.textAreaValue[:len(m.textAreaValue)-1]
				prompt := m.prompts[m.currentIndex].Prompt

				// Compare the input and get the highlighted input
				highlightedInput, _ := CompareAndHighlightInput(m.textAreaValue, prompt)
				m.textarea.Reset()
				m.textarea.Prompt = highlightedInput

			}
```

---

</SwmSnippet>

<SwmSnippet path="presentation/typingPage.go" line="168">

---

If the user click enter, if the text area has  characters we compare and highlight the typed content. We also calculate the typing accuracy and update our model.

&nbsp;

We then reset our text area, focus and set the text area to blank&nbsp;

If current index is less than the prompts, we continue to the next line and set it as the prompt, else we calculate the accuracy of the whole lesson and show the statistics at the end of the lesson.

```
		case tea.KeyEnter:
			input := m.textAreaValue

			if len(input) > 0 {
				prompt := m.prompts[m.currentIndex].Prompt
				highlightedInput, accuracy = CompareAndHighlightInput(input, prompt)
				m.prompt = append(m.prompt, fmt.Sprintf(" Input: %s (%.2f%% correct)\n", highlightedInput, accuracy))
				m.totalAccuracy += accuracy
			}

			m.textarea.Reset()
			m.textarea.Blur()
			m.viewport.GotoTop()
			m.textarea.Focus()

			// reset the text area value to blank
			m.textAreaValue = ""
			//Reintroduce the placeholder
			m.textarea.Placeholder = "Type the prompt"

			m.currentIndex++
			if m.currentIndex < len(m.prompts) {

				prompt := m.prompts[m.currentIndex].Prompt
				typingProgress := float64(len(input)) / float64(len(prompt))
				m.progress.Progress.SetPercent(typingProgress)

				m.prompt = append(m.prompt, " Prompt: "+m.promptStyle.Render(prompt))

				m.input = fmt.Sprintf("%s %s", m.input, prompt)
				m.viewport.SetContent(strings.Join(m.prompt, "\n"))

			} else {
				averageAccuracy := m.totalAccuracy / float64(len(m.prompts))
				database.WriteToDebugFile("m.input is while displaying ->", m.input)
				m.textarea.Prompt = " "
				m.textarea.Placeholder = ""
				m.prompt = append(m.prompt, m.resultsStyle.Render(typing.DisplayTypingSpeed(m.startTime, m.input, m.lesson, averageAccuracy)))
				m.viewport.SetContent(strings.Join(m.prompt, "\n"))
				return m, tea.Quit
			}
		}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="42">

---

Current index tracks our  position on the lesson contents list

```go
	currentIndex     int
	startTime        time.Time
	totalAccuracy    float64
	hasStartedTyping bool
	progress         progressBar.ProgressModel
}

func TypingPage(lesson models.Lesson) {
	p := tea.NewProgram(initialModel(lesson))

```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="211">

---

If it's any other key apart from the Backspace, Enter and escape, we update the text area as we validate the input  character against the prompt character. we show the user real time feedback through different colors.

```go
	default:
		// Validate while typing for real-time feedback
		input := m.textarea.Value()

		m.textAreaValue = m.textAreaValue + input

		if len(input) > 0 && m.currentIndex < len(m.prompts) {
			prompt := m.prompts[m.currentIndex].Prompt

			// Compare the input and get the highlighted input
			highlightedInput, _ := CompareAndHighlightInput(m.textAreaValue, prompt)
			m.textarea.Reset()
			m.textarea.Placeholder = ""
			m.textarea.Prompt = highlightedInput

		}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="228">

---

If the event is an error, we return the error

```go
	case errMsg:
		m.err = msg
		return m, nil
	}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="233">

---

We batch the update and return it here. Bubble tea can now update our view accordingly.

```go
	return m, tea.Batch(tiCmd, vpCmd)
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="236">

---

This is the view function that bubble uses to render our output on the terminal. we show the viewport, text area and progress bar. we also give users a hint that if they press esc they navigate back.

```go
func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
		m.progress.View(),
		" (Press esc to exit)",
	) + "\n"
}
```

---

</SwmSnippet>

<SwmSnippet path="/presentation/typingPage.go" line="245">

---

This function helps us show the user real time feedback by comparing and highlighting the input and the prompt. Red for error and Green for good. we also return the accuracy as we compare

```go
func CompareAndHighlightInput(input string, prompt string) (string, float64) {
	correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))   // Green for correct characters
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1")) // Red for incorrect characters

	var highlightedInput strings.Builder
	correctCount := 0

	for i := 0; i < len(input) && i < len(prompt); i++ {
		if input[i] == prompt[i] {
			highlightedInput.WriteString(correctStyle.Render(string(input[i])))
			correctCount++
		} else {
			highlightedInput.WriteString(incorrectStyle.Render(string(input[i])))
		}
	}

	accuracy := float64(correctCount) / float64(len(prompt)) * 100
	return highlightedInput.String(), accuracy
}
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
