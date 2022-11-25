package infrastructure

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func LoadEnvVariables() error {
	return godotenv.Load()
}

var dbConnection *gorm.DB

func DbConnection() (*gorm.DB, error) {
	var err error

	if dbConnection == nil {
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_NAME"),
		)
		dbConnection, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			dbConnection = nil
		}
	}

	return dbConnection, err
}
