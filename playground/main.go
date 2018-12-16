package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/koreset/timeslice/models"
	"time"
)

type TestSheet struct {
	gorm.Model
	Title string
	Date  time.Time `gorm:"type:date"`
}

func main() {
	db, err := gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/timeslice?parseTime=True&charset=utf8&loc=Local")
	defer db.Close()

	if err != nil {
		panic(err)
	}

	db.LogMode(true)
	populateData(db)

	var tsEntries []models.TimesheetEntry

	db.Where("entry_date between ? and ?", stringDate(time.Now().Add(-time.Hour*24*16)), stringDate(time.Now())).Find(&tsEntries)

	fmt.Println(tsEntries)

	db.Where("project_id = 1 and entry_date between ? and ?", stringDate(time.Now().Add(-time.Hour*24*16)), stringDate(time.Now())).Find(&tsEntries)

	fmt.Println(tsEntries)

}

func stringDate(t time.Time) string {
	return t.Format("2006-01-02")
}

func populateData(db *gorm.DB) {
	db.DropTableIfExists(&models.Client{}, &models.Employee{}, &models.Project{}, &models.TimesheetEntry{})
	db.AutoMigrate(&models.Client{}, &models.Employee{}, &models.Project{}, &models.TimesheetEntry{})

	var emp1 = models.Employee{
		FirstName: "Jome",
		LastName:  "Akpoduado",
		Email:     "jome@koreset.com",
		Password:  "wordpass15",
	}

	var client1 = models.Client{
		Name: "Koreset Solutions",
	}

	var project1 = models.Project{
		Name: "Tangent HR Revamp",
	}

	var project2 = models.Project{
		Name: "HOMEF Web Application",
	}
	db.Create(&emp1)
	db.Create(&client1)
	db.Create(&project1)
	db.Create(&project2)

	var ts2 = models.TimesheetEntry{
		Description: "Timesheet Entry Two",
		EntryDate:   time.Now().Add(-time.Hour * 24 * 7),
		EmployeeID:  emp1.ID,
		ClientID:    client1.ID,
		ProjectID:   project1.ID,
		Duration:    4.0,
	}

	var ts3 = models.TimesheetEntry{
		Description: "Timesheet Entry Three",
		EntryDate:   time.Now(),
		EmployeeID:  emp1.ID,
		ClientID:    client1.ID,
		ProjectID:   project1.ID,
		Duration:    4.0,
	}

	var ts4 = models.TimesheetEntry{
		Description: "Timesheet Entry Three",
		EntryDate:   time.Now(),
		EmployeeID:  emp1.ID,
		ClientID:    client1.ID,
		ProjectID:   project2.ID,
		Duration:    4.0,
	}

	var ts1 = models.TimesheetEntry{
		Description: "Timesheet Entry One",
		EntryDate:   time.Now().Add(-time.Hour * 24 * 15),
		EmployeeID:  emp1.ID,
		ClientID:    client1.ID,
		ProjectID:   project1.ID,
		Duration:    5.0,
	}

	db.Create(&ts1)
	db.Create(&ts2)
	db.Create(&ts3)
	db.Create(&ts4)

	fmt.Println(ts1)
	fmt.Println(ts2)
	fmt.Println(ts3)
	fmt.Println(ts4)
}
