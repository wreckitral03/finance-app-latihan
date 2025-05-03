package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Type      string    `json:"type" binding:"required,oneof=income expense"`
	Amount    float64   `json:"amount" binding:"required,gt=0"`
	Note      string    `json:"note"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// BeforeCreate will set a UUID rather than numeric ID.
func (t *Transaction) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}
