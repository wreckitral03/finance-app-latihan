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

	if err := db.DB.Create(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, transaction)
}

func GetTransactions(c *gin.Context) {
	var transactions []models.Transaction
	db.DB.Order("created_at desc").Find(&transactions)
	c.JSON(http.StatusOK, transactions)
}

func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction
	if err := db.DB.First(&transaction, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, transaction)
}

func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")
	var transaction models.Transaction

	if err := db.DB.First(&transaction, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Save(&transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, transaction)
}

func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	if err := db.DB.Delete(&models.Transaction{}, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted"})
}

func GetSummary(c *gin.Context) {
	type SummaryResult struct {
		Income  float64
		Expense float64
	}
	var result SummaryResult

	db.DB.Raw(`
		SELECT 
			SUM(CASE WHEN type = 'income' THEN amount ELSE 0 END) as income,
			SUM(CASE WHEN type = 'expense' THEN amount ELSE 0 END) as expense
		FROM transactions
	`).Scan(&result)

	balance := result.Income - result.Expense

	c.JSON(http.StatusOK, gin.H{
		"total_income":   result.Income,
		"total_expenses": result.Expense,
		"balance":        balance,
	})
}
