package api

import (
	"finance-app/internal/db"
	"finance-app/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.DB.Create(&transaction)
	c.JSON(http.StatusOK, transaction)
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	db.DB.Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func GetSummary(c *gin.Context) {
	var incomeTotal, expenseTotal float64

	db.DB.Model(&models.Transaction{}).Where("type = ?", "income").Select("SUM(amount)").Scan(&incomeTotal)
	db.DB.Model(&models.Transaction{}).Where("type = ?", "expense").Select("SUM(amount)").Scan(&expenseTotal)

	balance := incomeTotal - expenseTotal

	c.JSON(http.StatusOK, gin.H{
		"total_income":   incomeTotal,
		"total_expenses": expenseTotal,
		"balance":        balance,
	})
}
