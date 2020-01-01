package config

import (
	"log"
	"github.com/jinzhu/gorm"
)

// open connection to postgre
func DbConnection() *gorm.DB {
	db, err := gorm.Open("postgres",
		"user="+ USERNAME +" password="+ PASSWORD +" dbname="+ DB_NAME +" host=localhost port="+DB_PORT+" sslmode=disable")

	// debug openning connection
	if err != nil {
		log.Printf("Connection failed : %v", err)
	} else {
		log.Println("Connection success")
	}

	return db
}
