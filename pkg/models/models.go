package models

import (
	"github.com/jinzhu/gorm"
)

type Lesson struct {
	gorm.Model
	Title        string
	Content      []string `gorm:"-"`
	CurrentSpeed float64
	BestSpeed    float64
	Complete     bool
}
