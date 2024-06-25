package theme

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// User es la estructura para el usuario
type Theme struct {
	ThemeID string `json:"theme_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeTitle string `json:"theme_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// ValidateTheme validates the Theme struct
func ValidateTheme(theme *Theme) error {
	return validate.Struct(theme)
}
