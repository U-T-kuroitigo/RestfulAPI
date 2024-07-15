package theme

import (
	"time"

	"github.com/U-T-kuroitigo/RestfulAPI/tables/chapter"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Theme struct {
	ThemeID    string            `json:"theme_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeTitle string            `json:"theme_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	Chapter    []chapter.Chapter `gorm:"foreignKey:ThemeID;references:ThemeID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateTheme validates the Theme struct
func ValidateTheme(theme *Theme) error {
	return validate.Struct(theme)
}
