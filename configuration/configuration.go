package configuration

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" //driver para mysql
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

// Configuration crea un struct para el json
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	Database string
}


// GetConfiguration obtiene la configuración del json
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

// GetConnection obtiene una conexion a la bd
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.Database)
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user{})

	return db
}

type user struct {
	gorm.Model
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}
