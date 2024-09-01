package models

type Practice struct {
	ID      uint     `gorm:"primary_key;autoIncrement:true" json:"id"`
	Title   string   `json:"title"`
	Lessons []Lesson `gorm:"foreignKey:PracticeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lessons"`
}
