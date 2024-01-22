// This file contains types that are used in the repository layer.
package repository

import (
	"time"

	"github.com/google/uuid"
)

// Users ...
type Users struct {
	ID           uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	PhoneNumber  string    `json:"phone_number"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	SuccessLogin int64     `json:"success_login"`
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
}
