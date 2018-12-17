package models

import (
	"database/sql/driver"
	"github.com/jinzhu/gorm"
	"gopkg.in/gin-gonic/gin.v1/json"
	"time"
)

type TimesheetEntry struct {
	gorm.Model
	ClientID    uint      `gorm:"not null" json:"client_id"`
	EmployeeID  uint      `gorm:"not null" json:"employee_id"`
	ProjectID   uint      `gorm:"not null" json:"project_id"`
	Duration    float64   `json:"duration"`
	EntryDate   EntryDate `gorm:"column:entry_date;type:date" json:"entry_date" time:"2016-01-02"`
	Description string    `json:"description"`
}

type EntryDate struct {
	time.Time
}

func (ed EntryDate) Value() (driver.Value, error) {
	return ed.Format("2006-01-02"), nil
}

func (ed *EntryDate) Scan(value interface{}) error {
	*ed = EntryDate{value.(time.Time)}
	return nil
}

func (ed *EntryDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(ed.Format("2006-01-02"))
}
