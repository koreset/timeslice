package controllers

import (
	"github.com/gin-gonic/gin"
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
