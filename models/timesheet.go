package models

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
	"time"
)

type TimesheetEntry struct {
	gorm.Model
	ClientID    uint      `gorm:"not null" json:"client_id"`
	EmployeeID  uint      `gorm:"not null" json:"employee_id"`
	ProjectID   uint      `gorm:"not null" json:"project_id"`
	Duration    float64   `json:"duration"`
	EntryDate   EntryDate `gorm:"type:date" json:"entry_date" time:"2016-01-02"`
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

func (ed *EntryDate) UnmarshalJSON(b []byte) error {
	var dateString string
	if err := json.Unmarshal(b, &dateString); err != nil {
		return err
	}

	if t, err := time.Parse("2006-01-02", dateString); err != nil {
		return err
	} else {
		*ed = EntryDate{t}
		return nil
	}
}
