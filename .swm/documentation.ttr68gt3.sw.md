---
id: ttr68gt3
title: documentation
file_version: 1.1.3
app_version: 1.15.0
---

Entry point to our app. We start with a welcome screen
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
16     func main() {
```

<br/>

receives the sentences as a list, iterate through list to allow user to input sentences one after the other and at the end display user stats to the user.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
<!-- collapsed -->

```go
18     func TypingPractice(lessonData *models.Lesson) {
19     	fmt.Println("Try this:")
20     	time.Sleep(delay)
```

<br/>

Takes in duration and all words and compares them to calculate words per minute
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/utils/typingSpeed/typingSpeed.go
<!-- collapsed -->

```go
8      func CalculateTypingSpeed(sentence string, duration time.Duration) float64 {
9      
10     	words := strings.Fields(sentence)
11     	wordCount := len(words)
12     	seconds := duration.Seconds()
13     	return float64(wordCount) / (seconds / 60.0)
14     }
```

<br/>

Convert the sentences into runes for easy comparisons and manipulation/
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
<!-- collapsed -->

```go
55     	sentenceCharacters := []rune(sentence)
56     
57     	for {
```

<br/>

Set the exit flag to true and break
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
<!-- collapsed -->

```go
32     	exitPractice := false
```

<br/>

compare the last character of the input characters with the character at the same index of the sentence characters
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
<!-- collapsed -->

```go
82     		if lastCharacter == sentenceCharacters[len(inputCharacters)-1] {
83     			fmt.Print(string(lastCharacter))
84     		} else {
85     			fmt.Printf("^")
86     		}
```

<br/>

initialize the sqlite database
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
17     	database.InitializeDatabase()
18     	root := "lessons"
```

<br/>

Check if the current lesson is complete in the database. /
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
75     func lessonComplete(lessonTitle string, lessons []models.LessonDTO) bool {
76     	for _, lesson := range lessons {
77     		if lesson.Title == lessonTitle {
78     			return true
79     		}
80     	}
81     	return false
82     }
```

<br/>

Read sentences from the txt files and returns a list of the sentences
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
84     // readLinesFromFile reads the lines from a file and returns them as a string slice.
85     func readLinesFromFile(filePath string) ([]string, error) {
86     	file, err := os.Open(filePath)
87     	if err != nil {
88     		return nil, err
89     	}
90     	defer file.Close()
91     
92     	var lines []string
93     	scanner := bufio.NewScanner(file)
94     	for scanner.Scan() {
95     		lines = append(lines, scanner.Text())
96     	}
97     	if err := scanner.Err(); err != nil {
98     		return nil, err
99     	}
100    	return lines, nil
101    }
```

<br/>

This welcome screen for each lesson to allow user to enter the lesson or exit the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
12     func WelcomeScreen(lessonData *models.Lesson, hasExitedLesson *bool) {
13     	clear.ClearScreen()
14     	fmt.Printf("Welcome to lesson %s\n", lessonData.Title)
15     	fmt.Println("\nPress RETURN or SPACE to continue to typing practice. Press ESC to quit")
16     
17     	if err := keyboard.Open(); err != nil {
18     		panic(err)
19     	}
20     	defer func() {
21     		_ = keyboard.Close()
22     	}()
23     
24     	for {
25     		_, key, err := keyboard.GetKey()
26     		if err != nil {
27     			break
28     		}
29     
30     		if key == keyboard.KeySpace || key == keyboard.KeyEnter {
31     			err := keyboard.Close()
32     			if err != nil {
33     				break
34     			}
35     			clear.ClearScreen()
36     			typing.TypingPractice(lessonData)
37     
38     		}
39     
40     		if key == keyboard.KeyEsc {
41     			keyboard.Close()
42     			*hasExitedLesson = true
43     			fmt.Printf("Exiting lesson %s...\n", lessonData.Title)
44     			break
45     		}
46     	}
47     
48     }
49     
```

<br/>

Displays user typing speed and also stores the statistics to the database
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
90     func displayTypingSpeed(startTime time.Time, inputWords string, lessonTitle string) {
91     
92     	endTime := time.Now()
93     	duration := endTime.Sub(startTime)
94     	typingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)
95     	fmt.Printf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lessonTitle, typingSpeed)
96     	var lesson models.LessonDTO
97     	lesson.Speed = fmt.Sprintf("%.2f WPM", typingSpeed)
98     	lesson.Title = lessonTitle
99     	database.CompleteLesson(lesson)
100    }
101    
```

<br/>

initialize the sqlite database. create table lessons if it does not exist
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/initializeDatabase.go
```go
8      func InitializeDatabase() {
9      	db, err := sql.Open("sqlite3", "pecklin.db")
10     	if err != nil {
11     		panic(err)
12     	}
13     	defer db.Close()
14     
15     	_, err = db.Exec(`
16             CREATE TABLE IF NOT EXISTS lessons (
17                 id INTEGER PRIMARY KEY AUTOINCREMENT,
18                 lesson TEXT UNIQUE,
19                 speed TEXT
20             )
21         `)
22     	if err != nil {
23     		panic(err)
24     	}
25     	db.Close()
26     }
```

<br/>

This file was generated by Swimm. [Click here to view it in the app](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=/docs/ttr68gt3).
