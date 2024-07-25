package extra_history

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExtraHistory struct {
	ExtraHistoryID   string `json:"extra_history_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	UserID           string `json:"user_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ThemeID          string `json:"theme_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChapterID        string `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraSituationID string `json:"extra_situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	CorrectFlag      bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateExtraHistory(extra_history *ExtraHistory) error {
	return validate.Struct(extra_history)
}
