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

	db.AutoMigrate(&user{},&theme{},&chapter{},&situation{})

	return db
}

type user struct {
	ID        uint           `gorm:"primaryKey"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

type theme struct {
	ThemeID string `json:"theme_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeTitle string `json:"theme_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	Chapter []chapter `gorm:"foreignKey:ThemeID;references:ThemeID"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

type chapter struct {
	ChapterID string `json:"chapter_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ThemeID string `json:"theme_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	ChapterTitle string `json:"chapter_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	Situation []situation `gorm:"foreignKey:ChapterID;references:ChapterID"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

type situation struct {
	SituationID string `json:"situation_id" gorm:"type:varchar(255);primaryKey;not null" validate:"max=32"`
	ChapterID string `json:"chapter_id" gorm:"type:varchar(255);not null" validate:"max=32"`
	SituationTitle string `json:"situation_title" gorm:"type:varchar(255);not null" validate:"max=12"`
	SituationLevel uint `json:"situation_level"`
	CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}