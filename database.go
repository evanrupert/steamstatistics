package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Tag represents a app with a tag
type Tag struct {
	gorm.Model
	Appid uint32
	Tag   string
}

// RunDatabaseMigrations migrates all of the structs to the database
func RunDatabaseMigrations() {
	db, _ := OpenConnection()
	defer db.Close()

	db.AutoMigrate(&Tag{})
}

// InsertTagsIntoDatabase inserts the list of tags into the database
func InsertTagsIntoDatabase(appTags AppTags, db *gorm.DB) {
	for _, tag := range appTags.Tags {
		tag := Tag{Appid: appTags.AppID, Tag: tag}

		insertTag(&tag, db)
	}
}

func insertTag(tag *Tag, db *gorm.DB) {
	db.Create(tag)
}

// GetAppTagsFromDatabase returns the list of tags for an appID from the database
func GetAppTagsFromDatabase(appID uint32, db *gorm.DB) AppTags {
	var tags []Tag
	db.Where("appid = ?", appID).Find(&tags)

	tagStrings := make([]string, len(tags))
	for i, tag := range tags {
		tagStrings[i] = tag.Tag
	}

	return AppTags{AppID: appID, Tags: tagStrings}
}

// OpenConnection opens a database connection
func OpenConnection() (*gorm.DB, error) {
	connString := os.Getenv("DATABASE_CONNECTION_STRING")

	db, err := gorm.Open("postgres", connString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
