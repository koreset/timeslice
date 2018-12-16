package services

import (
	"github.com/koreset/timeslice/models"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M){
	code := m.Run()
	os.Exit(code)
}

var emp models.Employee


func TestEmployeeService_Create(t *testing.T) {
	setupDB()
	emp = models.Employee{
		FirstName: "Jome",
		LastName:  "Akpoduado",
		Email:     "jome@koreset.com",
		Password:  "wordpass15",
	}

	empService := EmployeeService{}
	err := empService.Create(&emp)

	assert.Nil(t, err)
	assert.Equal(t, 60, len(emp.PasswordHash))
	assert.NotZero(t, emp.ID)
	cleanDB()
}


func setupDB(){
	db.DropTableIfExists(&models.Employee{})
	db.AutoMigrate(&models.Employee{})
}

func cleanDB(){
	db.DropTableIfExists(&models.Employee{})
}
