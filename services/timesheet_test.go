package services

import (
	"github.com/koreset/timeslice/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var tss = TimesheetService{}

func TestTimesheetService_CreateEntry(t *testing.T) {
	var tse = models.TimesheetEntry{
		Description:"sample description",
		Date: time.Now(),
		EmployeeID:1,
		ProjectID:1,
		ClientID:1,
		Duration:8.0,
	}

	assert.NotNil(t, tss.CreateEntry(&tse))

}
