package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gopkg.in/gin-gonic/gin.v1/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var testsc = TimesheetController{}

func TestTimesheetController_GetTimesheetsForEmployee(t *testing.T) {
	payload := make(map[string]interface{})

	payload["employee_id"] = 1
	payload["start_date"] = "2018-12-01"
	payload["end_date"] = "2018-12-17"

	r := gin.Default()
	r.GET("/employee/:empID/timesheet", testsc.GetTimesheetsForEmployee)

	jsonPayoad, _ := json.Marshal(payload)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/employee/1/timesheet", strings.NewReader(string(jsonPayoad)) )

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	content, _ := ioutil.ReadAll(w.Body)

	assert.Contains(t, string(content), "Timesheet Entry One")
	assert.Contains(t, string(content), "Timesheet Entry Two")
	assert.Contains(t, string(content), "Timesheet Entry Three")
}

func TestTimesheetController_CreateTimesheetEntryForEmployee(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/employee/:empID/timesheet", testsc.CreateTimesheetEntryForEmployee)
	payload := make(map[string]interface{})

	payload["employee_id"] = 1
	payload["project_id"] = 2
	payload["client_id"] = 1
	payload["duration"] = 8.0
	payload["entry_date"] = "2018-12-13"
	payload["description"] = "Another work done on understanding golang even better"
	jsonPayload, _ := json.Marshal(payload)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("POST", "/employee/1/timesheet", strings.NewReader(string(jsonPayload)))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	content, _ := ioutil.ReadAll(w.Body)

	assert.Contains(t, content, `"Another work done on understanding golang even better"`)
}
