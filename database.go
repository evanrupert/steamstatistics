package main

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// AppStatusNoSpecialStatus indicates a normal app
	AppStatusNoSpecialStatus = iota
	// AppStatusDoesNotExist indicates that an app does not have a steam store page
	AppStatusDoesNotExist
)

// Tag represents a app with a tag
type Tag struct {
	gorm.Model
	Appid uint32
	Tag   string
}

// AppStatus represents specialty statuses of some steam applications
// 0 = Does not exist
type AppStatus struct {
	gorm.Model
	Appid      uint32 `gorm:"not null;unique"`
	StatusCode uint16
}

// RunDatabaseMigrations migrates all of the structs to the database
func RunDatabaseMigrations() {
	db, _ := OpenConnection()
	defer db.Close()

	db.AutoMigrate(&Tag{}, &AppStatus{})
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

// GetAppStatusCode retrieves an applications status code
func GetAppStatusCode(appID uint32, db *gorm.DB) uint16 {
	var status AppStatus
	db.Where("appid = ?", appID).First(&status)

	return status.StatusCode
}

// InsertAppStatusCode inserts an appid with a given status into the app_statuses table
func InsertAppStatusCode(appID uint32, statusCode uint16, db *gorm.DB) {
	appStatus := AppStatus{Appid: appID, StatusCode: statusCode}

	db.Create(&appStatus)
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
