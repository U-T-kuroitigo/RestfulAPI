package choice

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Choice struct {
	ChoiceID string `json:"choice_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ProblemID string `json:"problem_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChoiceText string `json:"choice_text" gorm:"not null"`
	ChoiceExplanation string `json:"choice_explanation" gorm:"not null"`
	CorrectFlag bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateChoice(choice *Choice) error {
	return validate.Struct(choice)
}
