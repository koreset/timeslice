package main

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)
import _ "github.com/jinzhu/gorm/dialects/mysql"

type YesNoEnum bool

const (
	Yes YesNoEnum = true
	No            = false
)

type MyTime struct {
	time.Time
}

type Customer struct {
	CustomerID int64 `gorm:"primary_key"`
	Active     YesNoEnum
	Created    MyTime `gorm:"type:date"`
}

func (mt *MyTime) Scan(value interface{}) error {
	fmt.Println("Value is: ", value)
	*mt = MyTime{time.Now()}
	return nil
}

func (mt MyTime) Value() (driver.Value, error) {
	val := mt.Format("2006-01-02")
	fmt.Println(val)
	return val, nil
}

func (yne *YesNoEnum) Scan(value interface{}) error {
	// if value is nil, false
	if value == nil {
		// set the value of the pointer yne to YesNoEnum(false)
		*yne = YesNoEnum(false)
		return nil
	}
	if bv, err := driver.Bool.ConvertValue(value); err == nil {
		// if this is a bool type
		if v, ok := bv.(bool); ok {
			// set the value of the pointer yne to YesNoEnum(v)
			*yne = YesNoEnum(v)
			return nil
		}
	}
	// otherwise, return an error
	return errors.New("failed to scan YesNoEnum")
}

func main() {
	db, err := gorm.Open("mysql", "root:wordpass15@tcp(localhost:3306)/timeslice?parseTime=True")
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.LogMode(true)
	db.DropTableIfExists(&Customer{})
	db.AutoMigrate(&Customer{})

	c := &Customer{
		Active:  Yes,
		Created: MyTime{time.Now()},
	}

	fmt.Println(db.Create(&c).Error)

	var found Customer
	db.First(&found)
	fmt.Println(found)
}
