package main

import (
	"finance-app/internal/api"
	"finance-app/internal/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()

	router := gin.Default()

	router.POST("/transactions", api.CreateTransaction)
	router.GET("/transactions", api.GetTransactions)
	router.GET("/summary", api.GetSummary)

	router.Run(":8080")
}
