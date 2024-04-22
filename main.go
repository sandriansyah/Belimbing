package main

import (
	"anama/controllers/balance"
	"anama/controllers/customer"
	"anama/controllers/payment"
	"anama/controllers/product"
	"anama/controllers/productCategory"
	"anama/controllers/source"
	"anama/controllers/transaction"
	"anama/models"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())

	models.ConnectDatabase()
	models.DB.AutoMigrate()

	r.GET("/api/products", product.GetProducts)
	r.GET("/api/products/:id", product.GetProduct)
	r.POST("/api/products/add", product.CreateProduct)
	r.PUT("/api/products/:id", product.UpdateProduct)
	r.DELETE("/api/products/:id", product.DeleteProduct)

	r.GET("/api/customers", customer.GetCustomers)
	r.GET("/api/customers/:id", customer.GetCustomer)
	r.POST("/api/customer", customer.CreateCustomer)
	r.PUT("/api/customers/:id", customer.UpdateCustomer)
	r.DELETE("/api/customers/:id", customer.DeleteCustomer)

	r.GET("/api/transactions", transaction.GetTransactions)
	r.GET("/api/transactions/:id", transaction.GetTransaction)
	r.GET("/api/transactions/customers/:id", transaction.GetTransactionsByCustomer)
	r.POST("/api/transactions/add", transaction.CreateTransaction)

	r.GET("/api/payments", payment.GetPayments)
	r.GET("/api/payments/customers/:id", payment.GetPaymentsByCustomer)
	r.GET("/api/payments/:id", payment.GetPayment)
	r.POST("/api/payments/add", payment.CreatePayment)

	r.GET("/api/sources", source.GetSources)

	r.GET("/api/product-categories", productCategory.GetProductCategory)

	r.GET("/api/balances/sources/:id", balance.GetBalances)

	r.Run()

}
