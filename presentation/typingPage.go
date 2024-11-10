package presentation

import (
	"errors"
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	"github.com/charmbracelet/bubbles/progress"
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

var (
	accuracy         float64
	highlightedInput string
)

func TypingPage(lesson models.Lesson) {
	p := tea.NewProgram(initialModel(lesson))

	if finalModel, err := p.Run(); err != nil {
		log.Fatal(err)
	} else if lessonModel, ok := finalModel.(model); ok && lessonModel.err == nil {
		time.Sleep(time.Second * 7)
		navigation.Navigator.Pop()
	}
}

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

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
		m.progress.View(),
		" (Press esc to exit)",
	) + "\n"
}
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
