package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	// gorm postgres dialect
	_ "github.com/jinzhu/gorm/dialects/postgres"
	mocket "github.com/selvatico/go-mocket"
)

// DB *grom.DB
type DB struct {
	*gorm.DB
}

// ConnectDB : connecting DB
func ConnectDB() (*DB, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_SSL_MODE"))

	db, err := gorm.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return &DB{db}, nil
}

// GetMockDB : returns a mock db for test purpose
func GetMockDB() (*DB, error) {

	mocket.Catcher.Register() // Safe register. Allowed multiple calls to save
	mocket.Catcher.Logging = true

	connectionString := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("DB_HOST"),
		viper.GetInt("DB_PORT"),
		viper.GetString("DB_USER"),
		viper.GetString("DB_NAME"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_SSL_MODE"))

	mockDB, err := gorm.Open(mocket.DriverName, connectionString)

	if err != nil {
		log.Panic(
			fmt.Sprintf("Error connecting to database: %v", err),
			"db.init",
		)
	}

	return &DB{mockDB}, nil
}
