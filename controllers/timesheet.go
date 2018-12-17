package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/models"
	"github.com/koreset/timeslice/services"
	"net/http"
)

var tsc = services.TimesheetService{}

type TimesheetController struct {
}

type RequestPayload struct {
	EmployeeID uint   `json:"employee_id"`
	StartDate  string `json:"start_date"`
	EndDate    string `json:"end_date"`
}

type TimesheetPayload struct {
	Description string
	EntryDate   string
	ProjectID   uint
	ClientID    uint
	EmployeeID  uint
	Duration    float64
}

func (ts *TimesheetController) GetTimesheetsForEmployee(c *gin.Context) {
	var payload RequestPayload
	c.ShouldBindJSON(&payload)

	entries, err := tsc.FindEntriesByDateRangeForEmployee(payload.StartDate, payload.EndDate, payload.EmployeeID)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, entries)
}

func (ts *TimesheetController) CreateTimesheetEntryForEmployee(c *gin.Context) {
	var timesheetEntry models.TimesheetEntry
	err := c.ShouldBindJSON(&timesheetEntry)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
	}
	err = tsc.CreateEntry(&timesheetEntry)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err})
		return
	}

	c.JSON(http.StatusCreated, timesheetEntry)
}
