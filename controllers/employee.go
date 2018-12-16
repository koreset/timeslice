package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/models"
	"github.com/koreset/timeslice/services"
	"net/http"
)

type EmployeeController struct {
}

func (ec *EmployeeController) Create(c *gin.Context) {
	var newEmployee models.Employee

	err := c.ShouldBind(&newEmployee)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	employeeService := services.EmployeeService{}


	employeeService.Create(&newEmployee)
	fmt.Println(newEmployee)
	c.JSON(http.StatusOK, newEmployee)
}
