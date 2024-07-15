package extra_situation

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ExtraSituation struct {
	ExtraSituationID    string `json:"extra_situation_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ChapterID           string `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraSituationTitle string `json:"extra_situation_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateExtraSituation(extra_situation *ExtraSituation) error {
	return validate.Struct(extra_situation)
}
