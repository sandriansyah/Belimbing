package transaction

import (
	"anama/endpoints/request"
	"anama/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetTransactions(c *gin.Context) {

	var Transactions []models.Transaction

	models.DB.Preload("Customer").Preload("Product").Preload("Source").Find(&Transactions)
	c.JSON(http.StatusOK, gin.H{"data": Transactions})

}

func GetTransactionsByCustomer(c *gin.Context) {

	customerId := c.Param("id")
	var Transactions []models.Transaction

	models.DB.Preload("Customer").Preload("Product").Preload("Source").Find(&Transactions, "customer_id = ?", customerId)

	c.JSON(http.StatusOK, gin.H{"data": Transactions})

}

func GetTransaction(c *gin.Context) {
	var creditTransaction models.Transaction

	id := c.Param("id")

	if err := models.DB.Where("id=?", id).First((&creditTransaction)).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"data": creditTransaction})

}

func CreateTransaction(c *gin.Context) {
	var transactionReq request.CreateTransaction

	if err := c.ShouldBindJSON(&transactionReq); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	transaction := models.Transaction{
		CustomerId:       transactionReq.CustomerId,
		ProductId:        transactionReq.ProductId,
		SourceId:         transactionReq.SourceId,
		DownPayment:      transactionReq.DownPayment,
		InstallmentValue: transactionReq.InstallmentValue,
		InstallmentCount: transactionReq.InstallmentCount,
		PurchasePrice:    transactionReq.PurchasePrice,
	}

	var balance models.Balance
	models.DB.Where("source_id = ?", transactionReq.SourceId).Order("number desc").First(&balance)

	var product models.Product
	models.DB.Where("id = ?", transactionReq.ProductId).First(&product)

	var customer models.Customer
	models.DB.Where("id = ?", transactionReq.CustomerId).First(&customer)

	newBalance := models.Balance{
		Date:        transaction.Date,
		SourceId:    balance.SourceId,
		CashOut:     transaction.PurchasePrice,
		Saldo:       balance.Saldo - transaction.PurchasePrice,
		Description: customer.Name + "-" + product.ProductName + "-" + product.Type,
	}

	if err := models.DB.Create(&newBalance).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "message": err.Error()})
		return
	}

	if err := models.DB.Create(&transaction).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction": "success"})
}
