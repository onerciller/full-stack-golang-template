package entity

import (
	"time"
)

// BaseModel contains common fields for all models
// @Description Base model with common fields
type BaseModel struct {
	// Record ID
	ID uint `json:"id" gorm:"primarykey" example:"1"`
	// Creation timestamp
	CreatedAt time.Time `json:"created_at" example:"2024-02-19T14:20:00Z"`
	// Last update timestamp
	UpdatedAt time.Time `json:"updated_at" example:"2024-02-19T14:20:00Z"`
	// Deletion timestamp, null if not deleted
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index" swaggertype:"string" example:"2024-02-19T14:20:00Z"`
}
