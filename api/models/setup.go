package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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
	csvData := readData(filePath)

	airEntries := []*AirQuality{}
	if err := gocsv.UnmarshalBytes(csvData, &airEntries); err != nil {
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

func isURL(test string) bool {
	if _, err := url.ParseRequestURI(test); err != nil {
		return false
	}
	return true
}

func readURL(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		log.Panic(err)
	}

	defer func() {
		err := response.Body.Close()
		if err != nil {
			log.Panic(err)
		}
	}()
	if response.StatusCode != http.StatusOK {
		log.Panicf("Response status error: %v", response.StatusCode)
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err)
	}
	return body
}

func readFile(filePath string) []byte {

	content, err := ioutil.ReadFile(filepath.Clean(filePath))
	if err != nil {
		log.Panic(err)
	}
	return content
}

func readData(source string) []byte {
	if isURL(source) {
		log.Print("Init file is a remote url")
		return readURL(source)
	}
	log.Print("Init file is local")
	return readFile(source)
}
