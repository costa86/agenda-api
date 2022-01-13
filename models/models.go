package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Profile string `binding:"required"`
	Name    string `binding:"required"`
}

type TimeSlot struct {
	gorm.Model
	Slot     time.Time `binding:"required"`
	PersonID uint
}
