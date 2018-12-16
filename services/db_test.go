package services

import (
	"github.com/koreset/timeslice/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

//This is a moot test of sorts. If all conditions are good for mysql, then this should always pass
func TestGetDB(t *testing.T) {
	db, _ := GetDB()

	assert.NotNil(t, db)
	assert.Nil(t, err)
}

func TestGetDBEmployeeTableSetup(t *testing.T){
	db, _ := GetDB()

	assert.True(t,db.HasTable(&models.Employee{}) )
}
