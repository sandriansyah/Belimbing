package balance

import (
	"anama/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetBalances(c *gin.Context) {

	sourceId := c.Param("id")
	var balances []models.Balance

	models.DB.Preload("Source").Where("source_id=?", sourceId).Find(&balances)
	c.JSON(http.StatusOK, gin.H{"data": balances})

}
