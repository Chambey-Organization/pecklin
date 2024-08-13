package models

type Lesson struct {
	ID           uint            `gorm:"primary_key;autoIncrement:true" json:"id"`
	PracticeID   uint            `json:"-"`
	Title        string          `json:"title"`
	Content      []LessonContent `gorm:"foreignKey:LessonID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"content"`
	Input        string          `gorm:"-"`
	CurrentSpeed float64
	BestSpeed    float64
	Complete     bool
}
