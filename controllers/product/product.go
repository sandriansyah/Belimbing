package product

import (
	"anama/endpoints/request"
	"anama/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetProducts(c *gin.Context) {

	var products []models.Product

	if err := models.DB.Preload("Category").Find(&products).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"products": products})

}

func GetProduct(c *gin.Context) {
	var product models.Product
	id := c.Param("id")

	if err := models.DB.Where("id=?", id).First(&product).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"products": product})

}

func CreateProduct(c *gin.Context) {

	var productReq request.CreteProduct
	if err := c.ShouldBindJSON(&productReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	product := models.Product{
		ProductName: productReq.ProductName,
		Type:        productReq.Type,
		Description: productReq.Description,
		CategoryId:  productReq.CategoryId,
	}

	if err := models.DB.Create(&product).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": "success"})
}

func UpdateProduct(c *gin.Context) {

	var product models.Product

	id := c.Param("id")
	if err := c.ShouldBindJSON(&product); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&product).Where("id =?", id).Updates(&product).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})

}
func DeleteProduct(c *gin.Context) {

	productId := c.Param("id")

	id, err := uuid.Parse(productId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID tidak valid"})
		return
	}

	if err := models.DB.Delete(&models.Product{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"data": "Gagal menghapus produk"})
	}

}
