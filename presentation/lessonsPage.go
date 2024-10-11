package presentation

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"log"
	"main.go/data/local/database"
)

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
