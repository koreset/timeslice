package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type TimesheetEntry struct {
	gorm.Model
	ClientID    uint `gorm:"not null"`
	EmployeeID  uint `gorm:"not null"`
	ProjectID   uint `gorm:"not null"`
	Duration    float64
	EntryDate   time.Time `gorm:"type:date"`
	Description string
}
