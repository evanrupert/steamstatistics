package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Tag represents a app with a tag
type Tag struct {
	gorm.Model
	Appid uint64
	Tag   string
}

// RunDatabaseMigrations migrates all of the structs to the database
func RunDatabaseMigrations() {
	db := OpenConnection()
	defer db.Close()

	db.AutoMigrate(&Tag{})
}

// InsertTag inserts the given tag into the given database connection
func InsertTag(tag Tag, db *gorm.DB) {
	db.Create(tag)
}

// OpenConnection opens a database connection
func OpenConnection() *gorm.DB {
	connString := os.Getenv("DATABASE_CONNECTION_STRING")

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		panic(err)
	}

	return db
}
