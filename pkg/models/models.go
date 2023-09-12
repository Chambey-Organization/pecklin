package models

type Lesson struct {
    Title   string
    Content []string
}

type LessonDTO struct{
 Title string `json:"title"`
 Speed string `json:"speed"`
}