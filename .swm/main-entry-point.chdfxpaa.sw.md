---
title: Main - (Entry point)
---
<SwmSnippet path="/main.go" line="15">

---

This is the entry point of our app. We initialize the database, fetch the latest lessons and initialize navigation with the first page being the main menu

```go
func main() {
	database.InitializeDatabase()

	m := loader.InitialModel()
	p := tea.NewProgram(m)
	go func() {
		err := remote.FetchPractices()
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
```

---

</SwmSnippet>

<SwmMeta version="3.0.0" repo-id="Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=" repo-name="pecklin"><sup>Powered by [Swimm](https://app.swimm.io/)</sup></SwmMeta>
