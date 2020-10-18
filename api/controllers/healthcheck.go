package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrmarble/cartodb/models"
)

// HealthCheck Checks if database is available
func HealthCheck(c *gin.Context) {
	sqlDB, err := models.DB.DB()
	if err != nil {
		log.Panic(err)
	}
	if err = sqlDB.Ping(); err != nil {
		log.Print(err)
		c.String(http.StatusInternalServerError, "KO")
	} else {
		c.String(http.StatusOK, "OK")
	}
}
