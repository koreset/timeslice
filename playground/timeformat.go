package main

import (
	"fmt"
	"time"
)

type Person struct {
	Name     string
	Age      uint
	Birthday time.Time
}

func main() {
	//var person = Person{
	//	Name: "Jome Akpoduado", Age: 45, Birthday: time.Now(),
	//}
	//
	//bytes, err := json.Marshal(&person)
	//if err != nil{
	//	panic(err)
	//}
	//
	//fmt.Println(string(bytes))

	value  := "2018-18-12"
	// Writing down the way the standard time would look like formatted our way
	layout := "2006-02-01"
	t, _ := time.Parse(layout, value)
	fmt.Println(t)

	t2 := time.Now()
	fmt.Println(t2.Format("2006-01-02"))
}
