package configuration

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// Configuration creates a struct for the json
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

// GetConfiguration gets the configuration from the json
func GetConfiguration() Configuration {
	var c Configuration
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// 環境変数を取得
	c.Server = os.Getenv("Server")
	c.Port = os.Getenv("Port")
	c.User = os.Getenv("User")
	c.Password = os.Getenv("Password")
	c.Database = os.Getenv("Database")

	return c
}

// GetConnection obtains a connection to the database
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user{}, &user_profile{}, &theme{}, &chapter{}, &situation{}, &problem{}, &choice{}, &extra_situation{}, &extra_problem{}, &extra_choice{}, &history{})

	return db
}

type user struct {
	UserID      string         `json:"user_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	MailAddress string         `json:"mail_address" gorm:"index:,unique,type:varchar(255);not null"`
	GmailID     string         `json:"gmail_id" gorm:"unique,type:varchar(255);not null;size:255"`
	UserProfile []user_profile `gorm:"foreignKey:UserID;references:UserID"`
	History     []history      `gorm:"foreignKey:UserID;references:UserID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type user_profile struct {
	UserID    string `json:"user_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	UserName  string `json:"user_name" gorm:"not null"`
	UserImg   string `json:"user_img" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type theme struct {
	ThemeID    string    `json:"theme_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeTitle string    `json:"theme_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	ThemeImg   string    `json:"theme_img" gorm:"not null"`
	Chapter    []chapter `gorm:"foreignKey:ThemeID;references:ThemeID"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type chapter struct {
	ChapterID      string            `json:"chapter_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeID        string            `json:"theme_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChapterTitle   string            `json:"chapter_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	ChapterImg     string            `json:"chapter_img" gorm:"not null"`
	Situation      []situation       `gorm:"foreignKey:ChapterID;references:ChapterID"`
	ExtraSituation []extra_situation `gorm:"foreignKey:ChapterID;references:ChapterID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type situation struct {
	SituationID    string    `json:"situation_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ChapterID      string    `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	SituationTitle string    `json:"situation_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	SituationLevel uint      `json:"situation_level"`
	Problem        []problem `gorm:"foreignKey:SituationID;references:SituationID"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type problem struct {
	ProblemID          string `json:"problem_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	SituationID        string `json:"situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ProblemTitle       string `json:"problem_title" gorm:"not null"`
	ProblemText        string `json:"problem_text" gorm:"not null"`
	ProblemExplanation string `json:"problem_explanation" gorm:"not null"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}

type choice struct {
	ChoiceID          string `json:"choice_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ProblemID         string `json:"problem_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChoiceText        string `json:"choice_text" gorm:"not null"`
	ChoiceExplanation string `json:"choice_explanation" gorm:"not null"`
	CorrectFlag       bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type extra_situation struct {
	ExtraSituationID    string          `json:"extra_situation_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ChapterID           string          `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraSituationTitle string          `json:"extra_situation_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	ExtraProblem        []extra_problem `gorm:"foreignKey:ExtraSituationID;references:ExtraSituationID"`
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt `gorm:"index"`
}

type extra_problem struct {
	ExtraProblemID          string         `json:"extra_problem_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ExtraSituationID        string         `json:"extra_situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraProblemTitle       string         `json:"extra_problem_title" gorm:"not null"`
	ExtraProblemText        string         `json:"extra_problem_text" gorm:"not null"`
	ExtraProblemExplanation string         `json:"extra_problem_explanation" gorm:"not null"`
	ExtraChoice             []extra_choice `gorm:"foreignKey:ExtraProblemID;references:ExtraProblemID"`
	CreatedAt               time.Time
	UpdatedAt               time.Time
	DeletedAt               gorm.DeletedAt `gorm:"index"`
}

type extra_choice struct {
	ExtraChoiceID          string `json:"extra_choice_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ExtraProblemID         string `json:"extra_problem_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ExtraChoiceText        string `json:"extra_choice_text" gorm:"not null"`
	ExtraChoiceExplanation string `json:"extra_choice_explanation" gorm:"not null"`
	CorrectFlag            bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt              time.Time
	UpdatedAt              time.Time
	DeletedAt              gorm.DeletedAt `gorm:"index"`
}

type history struct {
	HistoryID   string `json:"history_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	UserID      string `json:"user_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ThemeID     string `json:"theme_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChapterID   string `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	SituationID string `json:"situation_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	CorrectFlag bool   `json:"correct_flag" gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
