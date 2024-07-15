package problem

import (
	"time"

	"github.com/U-T-kuroitigo/RestfulAPI/tables/choice"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Problem struct {
	ProblemID string `json:"problem_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	SituationID string `json:"situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ProblemTitle string `json:"problem_title" gorm:"not null"`
	ProblemText string `json:"problem_text" gorm:"not null"`
	ProblemExplanation string `json:"problem_explanation" gorm:"not null"`
	Choice []choice.Choice `gorm:"foreignKey:ProblemID;references:ProblemID"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateProblem(problem *Problem) error {
	return validate.Struct(problem)
}
