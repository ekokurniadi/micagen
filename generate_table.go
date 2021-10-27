package main

import (
	"log"

	"gorm.io/gorm"
)

func GenerateTable(db *gorm.DB, model interface{}) (string, error) {
	err := db.AutoMigrate(&model)
	if err != nil {
		log.Fatal("Create Table is not complete")
	}
	message := "Successfully"
	return message, err

}
