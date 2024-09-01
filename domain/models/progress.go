package models

type Progress struct {
	Id           uint `gorm:"primary_key;autoIncrement:true" json:"id"`
	LessonID     uint `gorm:"uniqueIndex;constraint:OnDelete:CASCADE" json:"lesson_id"`
	CurrentSpeed float64
	BestSpeed    float64
	Accuracy     float64
	Complete     bool
}
