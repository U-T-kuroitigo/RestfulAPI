package chapter

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Chapter struct {
	ChapterID string `json:"chapter_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeID string `json:"theme_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChapterTitle string `json:"chapter_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateChapter(chapter *Chapter) error {
	return validate.Struct(chapter)
}
