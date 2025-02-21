package entity

import (
	"time"
)

// BaseModel contains common fields for all models
// @Description Base model with common fields
type BaseModel struct {
	// Record ID
	ID uint `json:"id" gorm:"primarykey"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at"`
	// Deletion timestamp, null if not deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
