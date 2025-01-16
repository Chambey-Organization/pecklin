package presentation

import (
	"errors"
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/controllers/progressBar"
	"main.go/pkg/controllers/typing"
	"strings"
	"time"
)

type (
	errMsg error
)

type model struct {
	viewport         viewport.Model
	prompt           string
	input            string
	textAreaValue    string
	textAreaInput    string
	placeHolder      string
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

func TypingPage(lesson models.Lesson) {
	p := tea.NewProgram(initialModel(lesson))

	if finalModel, err := p.Run(); err != nil {
		log.Fatal(err)
	} else if lessonModel, ok := finalModel.(model); ok && lessonModel.err == nil {
		navigation.Navigator.Pop()
		navigation.Navigator.Navigate(func() {
			LessonResultsPage(lesson.ID)
		})
	}
}

func initialModel(lesson models.Lesson) model {
	ta := textarea.New()
	ta.Focus()
	ta.Prompt = " > "
	ta.SetWidth(100)
	ta.SetHeight(1)
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()
	ta.ShowLineNumbers = false

	ta.KeyMap.InsertNewline.SetEnabled(false)

	var progressBarModel progressBar.ProgressModel

	if lesson.TimerCount != nil {
		progressBarModel = progressBar.NewProgressModel(lesson.TimerCount)
	}

	m := model{
		textarea:         ta,
		viewport:         viewport.New(100, 5),
		err:              nil,
		lesson:           &lesson,
		textAreaInput:    "",
		placeHolder:      "",
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
		if m.currentIndex < len(m.lesson.Content) {
			prompt := m.lesson.Content[m.currentIndex].Prompt
			m.input = prompt
			m.placeHolder = prompt
			m.prompt = " " + m.promptStyle.Render(prompt)
		}
		m.viewport.SetContent(m.lesson.Title)
		m.startTime = time.Now()
	} else {
		m.placeHolder = "No prompts available"
	}

	return m
}

func (m model) Init() tea.Cmd {
	return tea.Batch(
		textarea.Blink,
		m.progress.Init(),
	)
}

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

				highlightedInput, _, placeHolder := m.CompareAndHighlightInput(m.textAreaValue, prompt)
				m.textarea.Reset()
				m.textarea.Prompt = "> " + highlightedInput
				m.placeHolder = placeHolder
				m.textAreaInput = highlightedInput
				database.WriteToDebugFile("Placeholder", placeHolder)
				typingProgress := float64(len(m.input)) / float64(len(prompt))
				m.progress.Progress.SetPercent(typingProgress)
			}
			return m, nil
		case tea.KeyEnter:
			input := m.textAreaValue

			if len(input) > 0 {
				prompt := m.prompts[m.currentIndex].Prompt
				highlightedInput, accuracy, _ := m.CompareAndHighlightInput(input, prompt)
				m.prompt = fmt.Sprintf(" Input: %s (%.2f%% correct)\n", highlightedInput, accuracy)
				m.totalAccuracy += accuracy
			}

			m.textarea.Reset()
			m.textarea.Blur()
			m.viewport.GotoTop()
			m.textarea.Focus()
			m.textarea.Prompt = "> "
			m.textAreaValue = ""
			m.textAreaInput = ""
			m.placeHolder = ""

			m.currentIndex++
			if m.currentIndex < len(m.prompts) {
				prompt := m.prompts[m.currentIndex].Prompt
				m.placeHolder = prompt

				typingProgress := float64(len(input)) / float64(len(prompt))
				m.progress.Progress.SetPercent(typingProgress)
				m.prompt = " " + m.promptStyle.Render(prompt)
				m.input = fmt.Sprintf("%s %s", m.input, prompt)
			} else {

				if err := typing.SaveTypingSpeed(m.startTime, m.input, m.lesson, m.totalAccuracy); err != nil {
					m.err = err
					database.WriteToDebugFile("An error happened saving the lesson", err.Error())
					return m, tea.Quit
				}

				return m, tea.Quit
			}
		}
	default:
		input := m.textarea.Value()
		m.textAreaValue += input
		if len(input) > 0 && m.currentIndex < len(m.prompts) {
			prompt := m.prompts[m.currentIndex].Prompt
			highlightedInput, _, placeHolder := m.CompareAndHighlightInput(m.textAreaValue, prompt)
			m.textarea.Reset()
			m.placeHolder = placeHolder
			m.textAreaInput = highlightedInput
			m.textarea.Prompt = "> " + highlightedInput
			typingProgress := float64(len(m.input)) / float64(len(prompt))
			m.progress.Progress.SetPercent(typingProgress)
		}
		return m, nil
	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {

	greyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("240")) // Adjust color code as needed

	displayText := m.textAreaInput + greyStyle.Render(m.placeHolder)

	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s",
		m.viewport.View(),
		"> "+displayText, // Display the text with styles applied
		m.progress.View(),
		" (Press esc to exit)",
	) + "\n"
}

func (m model) CompareAndHighlightInput(input string, prompt string) (string, float64, string) {
	correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))

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

	remainingPrompt := ""
	if len(prompt) > len(input) {
		remainingPrompt = prompt[len(input):]
	}

	accuracy := float64(correctCount) / float64(len(prompt)) * 100
	return highlightedInput.String(), accuracy, remainingPrompt
}
