package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/controllers"
)

func initializeRoutes(r *gin.Engine) {
	empC := controllers.EmployeeController{}
	tsC := controllers.TimesheetController{}

	r.POST("/employee", empC.Create)
	r.GET("/employee/:id/timesheet", tsC.GetTimesheetsForEmployee)
	r.POST("/employee/:id/timesheet", tsC.CreateTimesheetEntryForEmployee)
}
