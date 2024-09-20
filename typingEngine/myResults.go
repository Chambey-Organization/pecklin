package typingEngine

func ReadMyResults() {
	/*results := database.GetResults()

	var options []huh.Option[string]
	var number = 1
	for _, result := range results {
		optionText := fmt.Sprintf("%d. %s", number, result.LessonID)
		options = append(options, huh.NewOption(optionText, strconv.Itoa(int(result.LessonID))))
		number++
	}

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title("Main Menu").Options(
				options...,
			).Value(&practiceId).Validate(func(str string) error {
				if practiceId == "" {
					err := fmt.Sprintf("Please select a lesson to continue")
					return errors.New(err)
				}
				return nil
			}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	*/

}
