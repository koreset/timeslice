package main

import (
	"github.com/gin-gonic/gin"
	"github.com/koreset/timeslice/services"
	"log"
)

func main() {
	err := services.InitializeDB()
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	initializeRoutes(r)

	log.Fatal(r.Run(":8080"))
}
