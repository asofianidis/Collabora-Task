package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase() (*gorm.DB, error) {
	dsn := "host=" + os.Getenv("DBHOST") + " user=" + os.Getenv("DBUSER") + " password=" + os.Getenv("DBPASSWORD") + " dbname=" + os.Getenv("DBNAME") + " port=" + os.Getenv("DBPORT")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	return db, err
}
