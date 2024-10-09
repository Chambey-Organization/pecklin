package presentation

import (
	"errors"
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
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
				Value(&selectedLessonIndex).
				Validate(func(i int) error {
					if i < 0 {
						return errors.New("select a lesson to continue")
					}
					return nil
				}),
		),
	)

	err = form.Run()
	if err != nil {
		log.Fatal(err)
	}

	lesson := practiceLessons[selectedLessonIndex]

	navigation.Navigator.Navigate(func() {
		TypingPage(lesson)
	},
	)
}
