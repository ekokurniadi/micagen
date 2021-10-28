package micagen

import (
	"log"
	"os"

	"github.com/ekokurniadi/micagen/core"
	"github.com/ekokurniadi/micagen/entity"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func example() {
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
		log.Fatal("Cannot connect to the database")
		return
	}

	// GenerateAll(db, &entity.User{})
	core.GenerateAll(db, &entity.Order{})
	// After file is generated please delete or comment function above if you don't won't to generate file using mica generator

	router := gin.Default()

	// handle cors
	router.Use(cors.Default())

	router.Run()

}
