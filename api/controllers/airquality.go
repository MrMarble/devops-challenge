package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrmarble/cartodb/models"
)

// GetAll returns all air data
func GetAll(c *gin.Context) {
	var airData []models.AirQuality
	models.DB.Debug().Find(&airData)
	c.JSON(http.StatusOK, gin.H{"data": airData})
}
