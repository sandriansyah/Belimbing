package customer

import (
	"anama/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetCustomers(c *gin.Context) {

	var customers []models.Customer
	models.DB.Find(&customers)
	c.JSON(http.StatusOK, gin.H{"data": customers})

}

func GetCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")

	if err := models.DB.Where("id=?", id).First(&customer).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": customer})

}

func CreateCustomer(c *gin.Context) {

	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	models.DB.Create(&customer)
	c.JSON(http.StatusOK, gin.H{"customer": customer})

}

func UpdateCustomer(c *gin.Context) {
	var customer models.Customer
	id := c.Param("id")
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if models.DB.Model(&customer).Where("id =?", id).Updates(&customer).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diupdate"})
}

func DeleteCustomer(c *gin.Context) {

	var customer models.Customer

	input := map[string]string{"id": "0"}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := strconv.ParseInt(input["id"], 10, 64)

	if models.DB.Delete(&customer, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Gagal menghapus product"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})

}
