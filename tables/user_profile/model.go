package user_profile

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type UserProfile struct {
	UserID    string `json:"user_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	UserName  string `json:"user_name" gorm:"not null"`
	UserImg   string `json:"user_img" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUserProfile(user_profile *UserProfile) error {
	return validate.Struct(user_profile)
}
