package models

import "time"

type Transaction struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Amount    float64   `json:"amount"`
	Type      string    `json:"type"` // "income" or "expense"
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
}
