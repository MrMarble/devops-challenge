package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/gocarina/gocsv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// DB represents the database with queries
var DB *gorm.DB

// ConnectDataBase opens the database
func ConnectDataBase(host string, port string, user string, password string, dbname string) {
	database, err := gorm.Open(postgres.Open(fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		host, port, user, password, dbname)),
		&gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to database!")
	}

	err = database.AutoMigrate(&AirQuality{})
	if err != nil {
		log.Panic("Failed to connect to database!")
	}
	DB = database
}

// LoadData reads a csv file into the database
func LoadData(filePath string) {
	file, err := os.OpenFile(filepath.Clean(filePath), os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Panic(err)
		}
	}()

	airEntries := []*AirQuality{}
	if err := gocsv.UnmarshalFile(file, &airEntries); err != nil {
		log.Panic(err)
	}
	for _, airQuality := range airEntries {
		// If data already exists ignore warnings
		DB.Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(airQuality)
	}
	DB.Commit()
	os.Exit(0)
}
