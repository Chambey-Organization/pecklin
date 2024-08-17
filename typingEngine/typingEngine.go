package typingEngine

import (
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"main.go/data/local/database"
	"main.go/domain/models"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/utils/clear"
	"strings"
	"time"
)

var (
	hasExitedLesson = false
	delay           = 3 * time.Second
	startTime       = time.Now()
)

func ReadPracticeLessons(practiceId uint) error {
	practiceLessons, _ := database.ReadPracticeLessons(practiceId)

	for index, lesson := range practiceLessons {

		if !hasExitedLesson {
			if index != 0 {
				time.Sleep(delay)
			}
			clear.ClearScreen()
			p := tea.NewProgram(initialModel(lesson))

			if _, err := p.Run(); err != nil {
				return err
			}
		} else {
			time.Sleep(delay)
			return errors.New("user exited the practice")

		}
	}

	return nil
}

func lessonComplete(lessonTitle string, lessons []models.Lesson) bool {
	for _, lesson := range lessons {
		if lesson.Title == lessonTitle {
			return true
		}
	}
	return false
}

type (
	errMsg error
)

type model struct {
	viewport      viewport.Model
	input         []string
	textarea      textarea.Model
	senderStyle   lipgloss.Style
	questionStyle lipgloss.Style
	titleStyle    lipgloss.Style
	err           error
	lesson        *models.Lesson
	prompts       []models.LessonContent
	currentIndex  int
	instructions  string
}

func initialModel(lesson models.Lesson) model {
	ta := textarea.New()
	ta.Placeholder = "Type the prompt"
	ta.Focus()

	ta.Prompt = "> "

	ta.SetWidth(100)
	ta.SetHeight(1)

	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	title := fmt.Sprintf("Welcome to lesson %s", lesson.Title)

	vp := viewport.New(100, 10)
	var titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#6361e4"))

	vp.SetContent(titleStyle.Render(title))

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:      ta,
		input:         []string{},
		viewport:      vp,
		err:           nil,
		lesson:        &lesson,
		prompts:       lesson.Content,
		questionStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("4")),
		titleStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("#6361e4")),
		currentIndex:  0,
		instructions:  "(Press Enter to continue, esc to exit)",
	}
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
			m.instructions = "(Press esc to exit)"

			answer := m.textarea.Value()

			if len(answer) > 0 {
				prompt := m.prompts[m.currentIndex].Prompt
				highlightedInput := compareAndHighlightInput(answer, prompt)
				m.input = append(m.input, m.senderStyle.Render(fmt.Sprintf("Input: %s", highlightedInput)))
			} else {
				startTime = time.Now()
			}
			m.textarea.Reset()
			m.viewport.GotoTop()

			prompt := m.prompts[m.currentIndex]
			m.input = append(m.input, m.questionStyle.Render("Prompt: ")+prompt.Prompt)
			m.lesson.Input = fmt.Sprintf(m.lesson.Input, prompt)

			if m.currentIndex < len(m.prompts) {
				m.currentIndex++
				m.viewport.SetContent(strings.Join(m.input, "\n"))
			} else {
				m.input = append(m.input, m.senderStyle.Render(typing.DisplayTypingSpeed(startTime, m.lesson.Input, m.lesson.Title)))
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
		m.instructions,
	) + "\n"
}
func compareAndHighlightInput(input string, prompt string) string {

	fmt.Printf("\n>input %s > prompt %s", input, prompt)
	correctStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	incorrectStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("1"))

	var highlightedInput strings.Builder

	inputLen := len(input)
	promptLen := len(prompt)

	for i := 0; i < inputLen; i++ {
		if inputLen == promptLen {
			if i < promptLen && input[i] == prompt[i] {
				highlightedInput.WriteString(correctStyle.Render(string(input[i])))
			} else {
				highlightedInput.WriteString(incorrectStyle.Render(string(input[i])))
			}
		}
	}

	return highlightedInput.String()
}
