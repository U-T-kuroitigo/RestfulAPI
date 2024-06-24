package user

import (
	"time"

	"gorm.io/gorm"
)

// User es la estructura para el usuario
type User struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}
