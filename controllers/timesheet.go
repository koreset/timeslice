package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TimesheetController struct {
}

func (ts *TimesheetController) GetTimesheetsForEmployee(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"error": "Not Yet Implemented"})
}
