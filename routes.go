package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/controllers"
)

func initializeRoutes(r *gin.Engine) {
	empC := controllers.EmployeeController{}
	tsC := controllers.TimesheetController{}

	r.POST("/employees/new", empC.Create)
	r.GET("/timesheets", tsC.GetTimesheetsForEmployee)
}
