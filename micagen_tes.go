package main

import (
	"log"
	"os"

	"github.com/ekokurniadi/micagen/entity"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	env := godotenv.Load()
	if env != nil {
		log.Fatal("Error loading .env file")
	}

	hostname := os.Getenv("HOST_NAME")
	user_host := os.Getenv("USER_HOST")
	pass_host := os.Getenv("PASS_HOST")
	database := os.Getenv("DATABASE")
	dsn := "" + user_host + ":" + pass_host + "@tcp(" + hostname + ")/" + database + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Cannot connect the database")
		return
	}

	GenerateTable(db, &entity.User{})
	CreateRepository(db, &entity.User{})

}
