package extra_problem

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExtraProblem struct {
	ExtraProblemID          string `json:"extra_problem_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ExtraSituationID        string `json:"extra_situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraProblemTitle       string `json:"extra_problem_title" gorm:"not null"`
	ExtraProblemText        string `json:"extra_problem_text" gorm:"not null"`
	ExtraProblemExplanation string `json:"extra_problem_explanation" gorm:"not null"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateExtraProblem(extra_problem *ExtraProblem) error {
	return validate.Struct(extra_problem)
}
