package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/controllers"
)

func initializeRoutes(r *gin.Engine) {
	empC := controllers.EmployeeController{}
	tsC := controllers.TimesheetController{}

	r.POST("/employee/new", empC.Create)
	r.GET("/employee/:empID/timesheet", tsC.GetTimesheetsForEmployee)
	r.POST("/employee/:empID/timesheet", tsC.CreateTimesheetEntryForEmployee)
}
