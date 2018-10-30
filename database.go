package main

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

const (
	// AppStatusNormalStatus indicates that an application is being stored normally
	AppStatusNormalStatus = iota
	// AppStatusDoesNotExist indicates that an app does not have a steam store page
	AppStatusDoesNotExist
)

// Tag represents a app with a tag
type Tag struct {
	gorm.Model
	AppID uint32
	Tag   string
}

// AppStatus represents specialty statuses of some steam applications
type AppStatus struct {
	gorm.Model
	AppID      uint32 `gorm:"not null;unique"`
	StatusCode uint16
}

// RunDatabaseMigrations migrates all of the structs to the database
func RunDatabaseMigrations() {
	db, _ := OpenConnection()
	defer db.Close()

	db.AutoMigrate(&Tag{}, &AppStatus{})
}

// StatusDatabaseWriter is a process that will write app statuse updates to the database
func StatusDatabaseWriter(statusChan chan AppStatus, stopChan chan bool, db *gorm.DB) {
	for true {
		select {
		case status := <- statusChan:
			db.Create(&status)
		case <- stopChan:
			fmt.Println("StatusDatabaseWriter process ending")
			break
		}
	}
}

func InsertTagsIntoDatabase(app App, db *gorm.DB) {
	for _, tag := range app.Tags {
		db.Create(&Tag{AppID: app.AppID, Tag: tag})
	}
}

// GetAppTagsFromDatabase returns the list of tags for an appID from the database
func GetAppTagsFromDatabase(appID uint32, db *gorm.DB) AppTags {
	var tags []Tag
	db.Where("app_id = ?", appID).Find(&tags)

	tagStrings := make([]string, len(tags))
	for i, tag := range tags {
		tagStrings[i] = tag.Tag
	}

	return AppTags{AppID: appID, Tags: tagStrings}
}

// GetAppStatusCode retrieves an applications status code
func GetAppStatusCode(appID uint32, db *gorm.DB) uint16 {
	var status AppStatus
	db.Where("app_id = ?", appID).First(&status)

	return status.StatusCode
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
