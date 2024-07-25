package user

import (
	"time"

	"github.com/U-T-kuroitigo/RestfulAPI/tables/extra_history"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/history"
	"github.com/U-T-kuroitigo/RestfulAPI/tables/user_profile"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type User struct {
	UserID       string                       `json:"user_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	MailAddress  string                       `json:"mail_address" gorm:"index:,unique,type:varchar(255);not null"`
	GmailID      string                       `json:"gmail_id" gorm:"unique,type:varchar(255);not null;size:255"`
	UserProfile  []user_profile.UserProfile   `gorm:"foreignKey:UserID;references:UserID"`
	History      []history.History            `gorm:"foreignKey:UserID;references:UserID"`
	ExtraHistory []extra_history.ExtraHistory `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateUser(user *User) error {
	return validate.Struct(user)
}
