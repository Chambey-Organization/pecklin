package models

type Progress struct {
	Id           uint   `gorm:"primary_key;autoIncrement:true" json:"id"`
	Lesson       Lesson `gorm:"foreignKey:PracticeID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"lessons"`
	CurrentSpeed float64
	BestSpeed    float64
	Accuracy     float64
	Complete     bool
}
