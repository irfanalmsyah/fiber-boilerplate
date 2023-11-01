package database

import (
	"fmt"
	"log"

	"github.com/irfanalmsyah/fiber-boilerplate/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
    var err error

    DB, err = gorm.Open(sqlite.Open("db.sqlite3"), &gorm.Config{})

    if err != nil {
        log.Fatalf("Could not connect to database: %v", err)
    }

    DB.AutoMigrate(&model.User{})
    fmt.Println("Connection Opened to Database")
}