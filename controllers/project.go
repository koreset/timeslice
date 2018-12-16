package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/models"
	"github.com/koreset/timeslice/services"
	"net/http"
)

type ProjectController struct {

}

func (ec *ProjectController) Create(c *gin.Context) {
	var newProject models.Project

	err := c.ShouldBind(&newProject)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	projectService := services.ProjectService{}


	projectService.Create(&newProject)
	fmt.Println(newProject)
	c.JSON(http.StatusOK, newProject)
}
