package models

type Lesson struct {
	Title   string
	Content []string
}

type LessonDTO struct {
	Title        string  `json:"title"`
	CurrentSpeed float64 `json:"currentSpeed"`
	BestSpeed    float64 `json:"bestSpeed"`
}
