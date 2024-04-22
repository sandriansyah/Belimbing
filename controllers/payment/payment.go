package payment

import (
	"anama/endpoints/request"
	"anama/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetPayments(c *gin.Context) {
	var payments []models.Payment

	if err := models.DB.Preload("Transaction").Preload("Customer").Preload("Product").Find(&payments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

func GetPaymentsByCustomer(c *gin.Context) {
	customerId := c.Param("id")

	var payments []models.Payment

	if err := models.DB.Preload("Transaction").Preload("Product").Preload("Customer").Where("customer_id=?", customerId).Find(&payments).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payments})
}

func GetPayment(c *gin.Context) {
	var creditPayment models.Payment

	id := c.Param("id")

	if err := models.DB.Where("id=?", id).First((&creditPayment)).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"creditPayment": creditPayment})

}

func CreatePayment(c *gin.Context) {
	var paymentReq request.Payment
	if err := c.ShouldBindJSON(&paymentReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	layout := "2006-01-02"
	parsedDate, err := time.Parse(layout, paymentReq.Date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid date format"})
		return
	}

	var transaction models.Transaction
	models.DB.Where("id=?", paymentReq.TransactionId).First(&transaction)

	payment := models.Payment{
		TransactionId: paymentReq.TransactionId,
		Amount:        paymentReq.Amount,
		InsatllmentTo: paymentReq.InsatllmentTo,
		Date:          parsedDate,
		CustomerId:    transaction.CustomerId,
		ProductId:     transaction.ProductId,
	}
	payment.Date = parsedDate

	var balance models.Balance
	models.DB.Where("source_id = ?", transaction.SourceId).Order("number desc").First(&balance)

	var product models.Product
	models.DB.Where("id = ?", transaction.ProductId).First(&product)

	var customer models.Customer
	models.DB.Where("id = ?", transaction.CustomerId).First(&customer)

	newBalance := models.Balance{
		Date:        payment.Date,
		SourceId:    balance.SourceId,
		CashIn:      int(payment.Amount),
		Saldo:       balance.Saldo + int(payment.Amount),
		Description: "Payment" + ":" + strconv.FormatFloat(float64(payment.InsatllmentTo), 'f', -1, 32) + customer.Name + "-" + product.ProductName + "-" + product.Type,
	}

	if err := models.DB.Create(&newBalance).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "message": err.Error()})
		return
	}

	if err := models.DB.Create(&payment).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": payment})
}
