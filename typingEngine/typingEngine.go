package typingEngine

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/utils"
	"strings"
	"time"
)

type (
	errMsg error
)

type model struct {
	viewport         viewport.Model
	input            []string
	textarea         textarea.Model
	senderStyle      lipgloss.Style
	questionStyle    lipgloss.Style
	titleStyle       lipgloss.Style
	err              error
	lesson           *models.Lesson
	prompts          []models.LessonContent
	currentIndex     int
	startTime        time.Time
	totalAccuracy    float64
	hasStartedTyping bool
}

var (
	hasExitedLesson  = false
	delay            = 5 * time.Second
	accuracy         float64
	highlightedInput string
)

func ReadPracticeLessons(practiceId uint) error {
	practiceLessons, err := database.ReadPracticeLessons(practiceId)
	if err != nil {
		return err
	}

	// Create a form for lesson selection
	var selectedLessonIndex int
	var options []huh.Option[int]
	for i, lesson := range practiceLessons {
		optionText := fmt.Sprintf("%d. %s", i+1, lesson.Title)
		options = append(options, huh.NewOption(optionText, i))
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[int]().
				Title("Which lesson do you want to practice?").
				Options(options...).
				Value(&selectedLessonIndex).
				Validate(func(i int) error {
					if i < 0 {
						return errors.New("please select a lesson to continue")
					}
					return nil
				}),
		),
	)

	err = form.Run()
	if err != nil {
		return err
	}

	utils.ClearScreen()
	lesson := practiceLessons[selectedLessonIndex]
	p := tea.NewProgram(initialModel(lesson))

	if _, err := p.Run(); err != nil {
		return err
	}

	return nil
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

	var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6361e4"))
	var input []string

	titleText := fmt.Sprintf(" Welcome to lesson %s", lesson.Title)
	input = append(input)
	title := fmt.Sprintf(titleStyle.Render(titleText))

	vp := viewport.New(100, 10)

	vp.SetContent(titleStyle.Render(title))

	ta.KeyMap.InsertNewline.SetEnabled(false)

	m := model{
		textarea:         ta,
		input:            []string{},
		viewport:         vp,
		err:              nil,
		totalAccuracy:    0,
		lesson:           &lesson,
		prompts:          lesson.Content,
		senderStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true),
		currentIndex:     0,
		hasStartedTyping: false,
	}

	if len(m.lesson.Content) > 0 {
		question := m.lesson.Content[m.currentIndex].Prompt
		m.input = append(m.input, m.senderStyle.Render(" Prompt: ")+question)
		m.startTime = time.Now()
		m.viewport.SetContent(strings.Join(m.input, "\n"))
	}

	return m
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			hasExitedLesson = true
			fmt.Println(m.textarea.Value())
			return m, tea.Quit
		case tea.KeyEnter:
			input := m.textarea.Value()

			if len(input) > 0 {
				prompt := m.prompts[m.currentIndex].Prompt
				highlightedInput, accuracy = CompareAndHighlightInput(input, prompt)
				m.input = append(m.input, m.senderStyle.Render(fmt.Sprintf(" Input: %s (%.2f%% correct)\n", highlightedInput, accuracy)))
				m.totalAccuracy += accuracy
			}

			m.textarea.Reset()
			m.viewport.GotoTop()
			m.currentIndex++
			if m.currentIndex < len(m.prompts) {
				prompt := m.prompts[m.currentIndex]
				m.input = append(m.input, m.titleStyle.Render(" Prompt: ")+prompt.Prompt)
				m.lesson.Input = fmt.Sprintf(m.lesson.Input, prompt)
				m.viewport.SetContent(strings.Join(m.input, "\n"))
			} else {
				averageAccuracy := m.totalAccuracy / float64(len(m.prompts))
				m.input = append(m.input, m.senderStyle.Render(typing.DisplayTypingSpeed(m.startTime, m.lesson.Input, m.lesson, averageAccuracy)))
				m.viewport.SetContent(strings.Join(m.input, "\n"))
				return m, tea.Quit
			}
		}

	case errMsg:
		m.err = msg
		return m, nil
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func (m model) View() string {
	return fmt.Sprintf(
		"%s\n\n%s\n\n%s",
		m.viewport.View(),
		m.textarea.View(),
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
