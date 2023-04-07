package database

import(
	"fmt"
	"log"
	"projek-pertama/sesi9/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host	="localhost"
	user	="postgres"
	password="admin"
	dbPort	="5432"
	dbName	="learning-gorm"
	db		*gorm.DB	
	err		error
)

func StartDB(){
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbName, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Product{})
}

func GetDB() *gorm.DB {
	return db
}