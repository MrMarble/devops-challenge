package main

import (
	"net/http"
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

	if len(os.Args) > 1 {
		models.LoadData(os.Args[1])
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	r.GET("/air_quality", controllers.GetAll)

	r.Run()
}
