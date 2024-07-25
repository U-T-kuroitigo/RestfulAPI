package extra_choice

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExtraChoice struct {
	ExtraChoiceID          string `json:"extra_choice_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ExtraProblemID         string `json:"extra_problem_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraChoiceText        string `json:"extra_choice_text" gorm:"not null"`
	ExtraChoiceExplanation string `json:"extra_choice_explanation" gorm:"not null"`
	CorrectFlag            bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateChoice(extra_choice *ExtraChoice) error {
	return validate.Struct(extra_choice)
}
