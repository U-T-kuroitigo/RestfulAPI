package situation

import (
	"time"

	"github.com/U-T-kuroitigo/RestfulAPI/tables/problem"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Situation struct {
	SituationID    string            `json:"situation_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ChapterID      string            `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	SituationTitle string            `json:"situation_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	SituationLevel uint              `json:"situation_level"`
	Problem        []problem.Problem `gorm:"foreignKey:SituationID;references:SituationID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateSituation(situation *Situation) error {
	return validate.Struct(situation)
}
