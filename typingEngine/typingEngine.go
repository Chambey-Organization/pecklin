package typingEngine

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	"main.go/pkg/controllers/typing"
	"main.go/pkg/models"
	"main.go/pkg/utils/clear"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var (
	hasExitedLesson bool = false
)

func ReadTextLessons(lessons []models.Lesson, exitErr *bool, lessonType string) error {

	return filepath.Walk(lessonType, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))

			if lessonComplete(fileNameWithoutExt, lessons) {
				return nil
			}
			fileContent, err := readLinesFromFile(path)
			if err != nil {
				return err
			}

			lessonData := models.Lesson{
				Title:   fileNameWithoutExt,
				Content: fileContent,
			}

			if !hasExitedLesson {
				clear.ClearScreen()
				p := tea.NewProgram(initialModel(lessonData))

				if _, err := p.Run(); err != nil {
					log.Fatal(err)
				}
			} else {
				*exitErr = true
				time.Sleep(1 * time.Second)
				return errors.New("user exited the lesson")
			}
		}
		return nil
	})
}

func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) != "" {
			lines = append(lines, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
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
	messages      []string
	textarea      textarea.Model
	senderStyle   lipgloss.Style
	questionStyle lipgloss.Style
	err           error
	lesson        *models.Lesson
	prompts       []string
	currentIndex  int
}

func initialModel(lesson models.Lesson) model {
	ta := textarea.New()
	ta.Placeholder = "Type the prompt"
	ta.Focus()

	ta.Prompt = "> "

	ta.SetWidth(100)
	ta.SetHeight(1)

	// Remove cursor line styling
	ta.FocusedStyle.CursorLine = lipgloss.NewStyle()

	ta.ShowLineNumbers = false

	title := fmt.Sprintf("Welcome to lesson %s\nPress Enter to continue", lesson.Title)

	vp := viewport.New(100, 10)
	vp.SetContent(title)

	ta.KeyMap.InsertNewline.SetEnabled(false)

	return model{
		textarea:      ta,
		messages:      []string{},
		viewport:      vp,
		senderStyle:   lipgloss.NewStyle().Foreground(lipgloss.Color("#00cc00")),
		err:           nil,
		lesson:        &lesson,
		questionStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("4")),
		prompts:       lesson.Content,
		currentIndex:  0,
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
			answer := m.textarea.Value()
			startTime := time.Now()

			if len(answer) > 0 {
				m.messages = append(m.messages, m.senderStyle.Render(fmt.Sprintf("Input: %s", answer)))
			}
			m.textarea.Reset()
			m.viewport.GotoBottom()

			if m.currentIndex < len(m.prompts) {
				prompt := m.prompts[m.currentIndex]
				m.messages = append(m.messages, m.questionStyle.Render("Prompt: ")+prompt)
				m.lesson.Input = fmt.Sprintf(m.lesson.Input, prompt)
				m.currentIndex++
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
			} else {
				m.messages = append(m.messages, m.senderStyle.Render(typing.DisplayTypingSpeed(startTime, m.lesson.Input, m.lesson.Title)))
				m.viewport.SetContent(strings.Join(m.messages, "\n"))
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
		"(esc to quit)",
	) + "\n"
}
