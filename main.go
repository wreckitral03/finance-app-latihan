package main

import (
	"finance-app/internal/api"
	"finance-app/internal/db"
	"finance-app/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize DB
	db.Connect()
	db.DB.AutoMigrate(&models.Transaction{})

	// Initialize Router
	r := gin.Default()

	// Routes
	r.POST("/transactions", api.CreateTransaction)
	r.GET("/transactions", api.GetTransactions)
	r.GET("/transactions/:id", api.GetTransactionByID)
	r.PUT("/transactions/:id", api.UpdateTransaction)
	r.DELETE("/transactions/:id", api.DeleteTransaction)
	r.GET("/summary", api.GetSummary)

	// Start server
	r.Run(":8080")
}
