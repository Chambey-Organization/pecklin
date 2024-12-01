package models

type LessonContent struct {
	ID       uint   `gorm:"primary_key;autoIncrement:true" json:"id"`
	LessonID uint   `json:"-"`
	Prompt   string `json:"prompt"`
	Active   bool   `json:"active"`
}
