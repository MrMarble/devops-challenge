package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mrmarble/cartodb/controllers"
	"github.com/mrmarble/cartodb/models"
)

func main() {

	models.ConnectDataBase(
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
	)

	// If there is an argument, call the migration
	if len(os.Args) > 1 {
		models.LoadData(os.Args[1])
	}

	r := gin.Default()

	r.GET("/", controllers.HealthCheck)

	r.GET("/air_quality", controllers.GetAll)

	err := r.Run()
	if err != nil {
		log.Panic(err)
	}
}
