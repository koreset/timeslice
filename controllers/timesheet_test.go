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

func TestTimesheetController_GetTimesheetsForEmployee(t *testing.T) {
	payload := make(map[string]interface{})

	payload["employee_id"] = 1
	payload["start_date"] = "2018-12-01"
	payload["end_date"] = "2018-12-16"

	r := gin.Default()
	tsc := TimesheetController{}
	r.GET("/timesheets", tsc.GetTimesheetsForEmployee)

	jsonPayoad, _ := json.Marshal(payload)
	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/timesheets", strings.NewReader(string(jsonPayoad)) )

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	content, _ := ioutil.ReadAll(w.Body)

	assert.Contains(t, string(content), "Timesheet Entry One")
	assert.Contains(t, string(content), "Timesheet Entry Two")
	assert.Contains(t, string(content), "Timesheet Entry Three")
}
