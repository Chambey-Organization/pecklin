---
id: kazxo8vb
title: Documentation
file_version: 1.1.3
app_version: 1.17.0
---

Initialize the database. We use GORM since it reduces boilerplate code. Create database named pecklin. You can access the file named pecklin.db on the root directory
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/initializeDatabase.go
```go
12     	var err error
13     	DB, err = gorm.Open("sqlite3", "pecklin.db")
14     	if err != nil {
15     		panic("Failed to connect to database")
16     	}
```

<br/>

create the lesson table from the models if it't exist
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/initializeDatabase.go
```go
19     	DB.AutoMigrate(&models.Lesson{})
```

<br/>

initialize the database. Helps us create the tables if they do not exist. It also creates a database connection to our local database
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
19     	database.InitializeDatabase()
```

<br/>

Read the completed lessons and store them in the completed lessons variable
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
21     	lessons := database.ReadCompletedLesson()
```

<br/>

These are all the lessons regardless of their completion status. Store them i allLessons variable
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
22     	allLessons := database.ReadAllLessons()
```

<br/>

Clear the screen so that the user can focus on the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
28     	clear.ClearScreen()
```

<br/>

Read the lessons in the lessons folder and if the lesson is incomplete, the user can start typing it. The lessons are displayed to the user in the same order as they are in the lessons folder
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
24     	err := ReadTextLessons(lessons, allLessons, &exitErr)
```

<br/>

After completing the lessons without exiting, the user is given another menu to either redo, see the completed lessons stats or exit the typing practice.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
29     	fmt.Println("\n Congratulations! You have completed all the lessons \n \nPress RETURN to redo the typing practice, SPACE to view lesson stats and ESC to quit")
30     	if err := keyboard.Open(); err != nil {
31     		panic(err)
32     	}
33     	defer func() {
34     		_ = keyboard.Close()
35     	}()
36     
37     	for {
38     		_, key, err := keyboard.GetKey()
39     		if err != nil {
40     			break
41     		}
```

<br/>

On pressing ENTER, the user is able to redo the lessons. here we set the status of all lessons as incomplete and the user starts from lesson one
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
43     		if key == keyboard.KeyEnter {
44     			err := keyboard.Close()
45     			if err != nil {
46     				break
47     			}
48     			database.RedoLessons()
49     			lessons = database.ReadCompletedLesson()
50     
51     			err = ReadTextLessons(lessons, allLessons, &exitErr)
52     			if exitErr {
53     				return
54     			}
55     			if err != nil {
56     				return
57     			}
58     		}
59     
```

<br/>

If the user presses SPACE on the keyboard, we read the lessons and display them on this for loop. typing speed is shown as the best speed on that lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
60     		if key == keyboard.KeySpace {
61     			for _, lesson := range allLessons {
62     				fmt.Printf("\nLesson Title: %s\n", lesson.Title)
63     				fmt.Printf("Typing Speed: %.2f WPM\n", lesson.BestSpeed)
64     				fmt.Println("---------------------------------")
65     			}
66     		}
```

<br/>

On ESC, break the loop which exits the whole typing lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
67     		if key == keyboard.KeyEsc {
68     			break
69     		}
```

<br/>

This is a variable which is passed as a pointer to other functions and it tracks if the user exits the typing practice at any point
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
20     	var exitErr bool
```

<br/>

Define the root folder containing the lessons in the .txt format
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
78     	root := "lessons"
```

<br/>

Read the lessons data as contained in the root folder. This will loop all files in the provided folder
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
80     	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
81     		if err != nil {
```

<br/>

We are only interested in files which are not directories. return nil if the file is a directory
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
87     		if !info.IsDir() {
```

<br/>

Extract the file name without extension in order to use it as the lesson title which will make the naming consistent throughout the code
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
88     			fileNameWithoutExt := strings.TrimSuffix(info.Name(), filepath.Ext(info.Name()))
89     
```

<br/>

Skip the lesson if it is complete
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
90     			if lessonComplete(fileNameWithoutExt, lessons) {
91     				return nil
92     			}
```

<br/>

If the lesson is not complete, read the file content and store the sentences in the fileContent variable
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
93     			fileContent, err := readLinesFromFile(path)
```

<br/>

Create the lesson data which is a struct that holds out data, lesson name and the content which is a list of sentences in the file
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
98     			lessonData := models.Lesson{
99     				Title:   fileNameWithoutExt,
100    				Content: fileContent,
101    			}
```

<br/>

Redirect user to the welcome screen of the lesson, pass the lesson data and has exited pointer to notify us if the user exits the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
103    			welcome.WelcomeScreen(&lessonData, &hasExitedLesson)
```

<br/>

If the user exits the lesson return an error which will trigger the end of the typing practice. Else delay the next lesson by 3 seconds so that the user to read his stats for the completed lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
106    			if hasExitedLesson {
107    				*exitErr = true
108    				return errors.New("user exited the lesson")
109    			} else {
110    				time.Sleep(3 * time.Second)
111    			}
```

<br/>

Go through the completed lessons and see if the lesson is complete. if complete return true. else false. helps in skipping the complete lessons
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
118    func lessonComplete(lessonTitle string, lessons []models.Lesson) bool {
119    	for _, lesson := range lessons {
120    		if lesson.Title == lessonTitle {
121    			return true
122    		}
123    	}
124    	return false
125    }
```

<br/>

Use scanner to read the sentences in the provided file and return a slice of sentences contained in the file, if an error is encountered while reading, return the error too.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 main.go
```go
127    func readLinesFromFile(filePath string) ([]string, error) {
128    	file, err := os.Open(filePath)
129    	if err != nil {
130    		return nil, err
131    	}
132    	defer file.Close()
133    
134    	var lines []string
135    	scanner := bufio.NewScanner(file)
136    	for scanner.Scan() {
137    		lines = append(lines, scanner.Text())
138    	}
139    	if err := scanner.Err(); err != nil {
140    		return nil, err
141    	}
142    	return lines, nil
143    }
```

<br/>

Clear the screen for the user to focus on the lesson without distractions from previous texts on the screen
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
13     	clear.ClearScreen()
```

<br/>

Welcome the user to the lesson and give them the menu, Press ENTER or SPACE to enter the lesson or ESC to exit the lesson.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
14     	fmt.Printf("Welcome to lesson %s\n", lessonData.Title)
15     	fmt.Println("\nPress RETURN or SPACE to continue to typing practice. Press ESC to quit")
16     
```

<br/>

Open the keyboard to start listening to keystrokes.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
17     	if err := keyboard.Open(); err != nil {
18     		panic(err)
19     	}
20     	defer func() {
21     		_ = keyboard.Close()
22     	}()
```

<br/>

On pressing SPACE or ENTER, clear the screen once again and redirect them to a typing practice page
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
29     
30     		if key == keyboard.KeySpace || key == keyboard.KeyEnter {
31     			err := keyboard.Close()
32     			if err != nil {
33     				break
34     			}
35     			clear.ClearScreen()
36     			typing.TypingPractice(lessonData, hasExitedLesson)
37     
38     		}
```

<br/>

On ESC, close the screen and set the has exited pointer to true. Also break from this loop
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/welcome/welcome.go
```go
40     		if key == keyboard.KeyEsc {
41     			keyboard.Close()
42     			*hasExitedLesson = true
43     			fmt.Printf("Exiting lesson %s...\n", lessonData.Title)
44     			break
45     		}
```

<br/>

This is the beginning of the lesson. Also delay them for one second so that they can get ready to start typing
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
18     	fmt.Println("Try this:")
19     	time.Sleep(delay)
```

<br/>

Initialize this as empty string. it will be used to append all words the user enters. very useful in calculating WPM
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
21     	inputWords := ""
22     
```

<br/>

Open the keyboard once again to listen to user input for the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
23     	if err := keyboard.Open(); err != nil {
24     		panic(err)
25     	}
26     	defer func() {
27     		_ = keyboard.Close()
28     	}()
```

<br/>

Initialize the start time so that we can use it in calculating time taken while on this practice
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
30     	startTime := time.Now()
```

<br/>

Initialize exit practice as false in order to track when the user exits the typing of this lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
31     	exitPractice := false
32     
```

<br/>

Display the sentence the user should type
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
34     		fmt.Printf("\n\n%s\n", sentence)
```

<br/>

Loop the lesson content and serve the user the sentences one after the other after they complete typing
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
33     	for _, sentence := range lessonData.Content {
34     		fmt.Printf("\n\n%s\n", sentence)
35     
36     		inputWords, exitPractice = handleTypingInput(sentence, inputWords)
37     
38     		if exitPractice {
39     			*hasExitedLesson = true
40     			break
41     		}
42     	}
```

<br/>

If the user did not exit the practice, display their statistics on the lesson they have just completed
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
44     	if !exitPractice {
45     		displayTypingSpeed(startTime, inputWords, lessonData.Title)
46     	}
```

<br/>

If the user exits the practice, exit the loop and set the hasexited pointer to true
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
38     		if exitPractice {
39     			*hasExitedLesson = true
40     			break
41     		}
```

<br/>

Initialize sentence characters into a slice of runes too
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
53     	sentenceCharacters := []rune(sentence)
54     
```

<br/>

Initialize the input characters slice of runes. useful in the comparison of sentence and input character if it is wrong or right
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
51     	var inputCharacters []rune
52     
```

<br/>

Listen to input characters and the key the user presses
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
56     		char, key, err := keyboard.GetKey()
57     		if err != nil {
58     			break
59     		}
```

<br/>

If the key pressed is ENTER, means the user is completed the sentence and wants to go to next sentence, break this loop.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
61     		if key == keyboard.KeyEnter {
62     			break
```

<br/>

If the key is escape, user wants to exit the lesson, break the loop and return has exited as true
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
63     		} else if key == keyboard.KeyEsc {
64     			fmt.Printf("\n\nExiting lesson ...\n")
65     			return inputWords, true
66     		} else if key == keyboard.KeySpace {
```

<br/>

If the key is space, append space to the input characters variable and space rune to the input characters slice
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
66     		} else if key == keyboard.KeySpace {
67     			inputWords += " "
68     			inputCharacters = append(inputCharacters, ' ')
69     		} else {
```

<br/>

Else just append the character to the input character and a string of the rune to input words
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
69     		} else {
70     			inputWords += string(char)
71     			inputCharacters = append(inputCharacters, char)
72     		}
```

<br/>

Break the loop if the length of input characters runes is more than the sentences runes
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
74     		if len(inputCharacters) > len(sentenceCharacters) {
75     			break
76     		}
```

<br/>

Access the last character of the input so as to use it to display back to the user
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
78     		lastCharacter := inputCharacters[len(inputCharacters)-1]
```

<br/>

Compare the last character rune and the character of the same index in the sentences runes. If they match the user entered the correct character, display it as a string, else display an error "^"
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
80     		if lastCharacter == sentenceCharacters[len(inputCharacters)-1] {
81     			fmt.Print(string(lastCharacter))
82     		} else {
83     			fmt.Printf("^")
84     		}
```

<br/>

Return user input words and false if the user did not exit the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
87     	return inputWords, false
```

<br/>

Get the current time and set it as end time for the lesson
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
92     	endTime := time.Now()
```

<br/>

Get the duration and calculate the typing speed
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
93     	duration := endTime.Sub(startTime)
94     	currentTypingSpeed := typingSpeed.CalculateTypingSpeed(inputWords, duration)
```

<br/>

Display the typing speed to the user
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
95     	fmt.Printf("\n\nCongratulations! You have completed lesson %s\nYour typing speed is: %.2f WPM\n", lessonTitle, currentTypingSpeed)
```

<br/>

Create a lessons struct and populate it with data then save it as complete in the database
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/controllers/typing/typing.go
```go
96     	var lesson models.Lesson
97     	lesson.CurrentSpeed = currentTypingSpeed
98     	lesson.BestSpeed = currentTypingSpeed
99     	lesson.Title = lessonTitle
100    	lesson.Complete = true
101    	database.CompleteLesson(lesson)
```

<br/>

Create an existing lessons struct to get the existing lesson before we can make changes to the database
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
9      	var existingLesson models.Lesson
```

<br/>

Get the lesson if it exists and append the data to the existing lesson struct.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
10     	DB.Where("title = ?", lesson.Title).First(&existingLesson)
11     
```

<br/>

If the lesson exists update current speed and best speed if the current speed is greater than best typing.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
13     		existingLesson.CurrentSpeed = lesson.CurrentSpeed
14     
15     		if lesson.CurrentSpeed > existingLesson.BestSpeed {
16     			existingLesson.BestSpeed = lesson.CurrentSpeed
17     		}
18     
19     		existingLesson.Complete = lesson.Complete
20     		DB.Save(&existingLesson)
```

<br/>

If the lesson doesn't exist, create it.
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
22     		DB.Create(&lesson)
```

<br/>

Convert all complete lessons to incomplete
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
26     func RedoLessons() {
27     	DB.Model(&models.Lesson{}).Where("complete = ?", true).Update("complete", false)
28     }
```

<br/>

Return a slice of all complete lessons
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
30     func ReadCompletedLesson() []models.Lesson {
31     	var lessons []models.Lesson
32     	DB.Where("complete = ?", true).Find(&lessons)
33     
34     	return lessons
35     }
```

<br/>

Returns a slice of all lessons
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 database/dao.go
```go
36     func ReadAllLessons() []models.Lesson {
37     	var lessons []models.Lesson
38     	DB.Find(&lessons)
39     
40     	return lessons
41     }
```

<br/>

initialize clear command for all os
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/utils/clear/clearScreen.go
```go
13     	clear["linux"] = func() {
14     		cmd := exec.Command("clear") //Linux
15     		cmd.Stdout = os.Stdout
16     		cmd.Run()
17     	}
18     	clear["windows"] = func() {
19     		cmd := exec.Command("cmd", "/c", "cls") //Windows
20     		cmd.Stdout = os.Stdout
21     		cmd.Run()
22     	}
23     	clear["darwin"] = func() {
24     		cmd := exec.Command("clear") // Darwin (macOS)
25     		cmd.Stdout = os.Stdout
26     		cmd.Run()
27     	}
```

<br/>

Execute the clear command
<!-- NOTE-swimm-snippet: the lines below link your snippet to Swimm -->
### 📄 pkg/utils/clear/clearScreen.go
```go
31     	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
32     	if ok {
33     		value() //we execute it
34     	}
```

<br/>

This file was generated by Swimm. [Click here to view it in the app](https://app.swimm.io/repos/Z2l0aHViJTNBJTNBcGVja2xpbiUzQSUzQWNoYW1iZXk=/docs/kazxo8vb).
