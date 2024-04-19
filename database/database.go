package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(hostname string, port string, database string, username string, password string) error {
	var err error
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", hostname, port, username, password, database)
	DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	return err
}

func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
