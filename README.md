# micagen
Rest Api Generator for Golang Programming Language

[![Go Reference](https://pkg.go.dev/badge/github.com/ekokurniadi/micagen.svg)](https://pkg.go.dev/github.com/ekokurniadi/micagen)

[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=ekokurniadi&repo=micagen&theme=radical&show_icons=true)](https://github.com/ekokurniadi/micagen)


### How to install
```sh
go get -u github.com/ekokurniadi/micagen
```

### If you are using MySql DB

 - Please run this command on your project terminal
```sh
go get -u gorm.io/driver/mysql
```

### If you are using PostgreSql DB
 - Please run this command on your project terminal
```sh
go get -u gorm.io/driver/postgres
```

### Then import the package

```go
import "github.com/ekokurniadi/micagen"
```
### Important

- Create folder with name entity on root directory project
- Create file on entity folder and create struct
- Example :  create 1 file with name customer.go on entity folder
```go
package entity

import "time"

type Customer struct {
	ID        int
	Name      string
	Address   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
```
### This case i want to migrate struct customer to my database and generate crud rest api using micagen


### Example project using micagen and MySql db
```go
package main

import (
	"log"
	"tesss/entity"

	"github.com/ekokurniadi/micagen"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/mica_generator?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
       
        //call micagen package
	gen := micagen.Micagen{}
        //automigrate struct of customer and generate crud rest api for customer
	gen.GenerateAll(db, &entity.Customer{})

}
```
### Example project using micagen and PostgreSql
```go
package main

import (
	"log"
	"tesss/entity"

	"github.com/ekokurniadi/micagen"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){

        env := godotenv.Load()
	if env != nil {
	   log.Fatal("Error loading .env file")
	 }

	host := os.Getenv("DB_HOST")
	userHost := os.Getenv("DB_USER")
	userPass := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_DATABASE")
	databasePort := os.Getenv("DB_PORT")

	dsn := "host=" + host + " user=" + userHost + " password=" + userPass + " dbname=" + databaseName + " port=" + databasePort + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println("Database Connected")

        //call micagen package
	gen := micagen.Micagen{}
        //automigrate struct of customer and generate crud rest api for customer
	gen.GenerateAll(db, &entity.Customer{})

}
```
https://github.com/ekokurniadi/micagen-example.git
