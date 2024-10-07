package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Age       int
	CreatedAt time.Time
}

func main() {
	db, err := gorm.Open(sqlite.Open("gfgTest.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	fmt.Println("Db got connected")
	db.AutoMigrate((&User{}))
}
