package productCategory

import (
	"anama/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetProductCategory(c *gin.Context) {
	var productCategories []models.Product_category

	models.DB.Find(&productCategories)
	c.JSON(http.StatusOK, gin.H{"data": productCategories})

}
