package main

import (
	"fmt"
	"github.com/CharlesMuchogo/GoNavigation/navigation"
	tea "github.com/charmbracelet/bubbletea"
	"main.go/data/local/database"
	"main.go/data/remote"
	"main.go/pkg/controllers/loader"
	"main.go/pkg/utils"
	"main.go/presentation"
	"os"
	"time"
)

func main() {
	database.InitializeDatabase()

	m := loader.InitialModel()
	p := tea.NewProgram(m)
	go func() {
		err := remote.FetchPractices()
		time.Sleep(5 * time.Second)
		if err != nil {
			p.Send(loader.DataLoadedMsg{})
			return
		}

		p.Send(loader.DataLoadedMsg{})
		utils.ClearScreen()
	}()

	if _, err := p.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	navigation.InitialRoute(func() {
		presentation.MainMenu()
	})
}
