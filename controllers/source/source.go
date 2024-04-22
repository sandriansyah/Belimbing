package source

import (
	"anama/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetSources(c *gin.Context) {
	var sources []models.Source

	models.DB.Find(&sources)
	c.JSON(http.StatusOK, gin.H{"data": sources})
}
