package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

func GetAllUserAppTags(steamID string) ([]App, error) {
	db, err := OpenConnection()
	if err != nil {
		return nil, err
	}

	apps, err := GetUserApps(steamID)
	if err != nil {
		return nil, err
	}

	appTagsArray := make([]App, len(apps))

	appChan := make(chan App)
	statusChan := make(chan AppStatus)
	stopChan := make(chan bool)

	go StatusDatabaseWriter(statusChan, stopChan, db)

	i := 0

	for _, app := range apps {
		appTags := GetAppTagsFromDatabase(app.AppID, db)
		if needsToGatherData(appTags, db) {
			go getAppTags(app, appChan, statusChan)
			i++
		} else {
			appTagsArray = append(appTagsArray, createApp(appTags, app.Playtime))
		}
	}

	for i > 0 {
		app := <-appChan
		InsertTagsIntoDatabase(app, db)
		appTagsArray = append(appTagsArray, app)
		i--
	}

	stopChan <- true

	return appTagsArray, nil
}

func needsToGatherData(appTags AppTags, db *gorm.DB) bool {
	return len(appTags.Tags) <= 0 && GetAppStatusCode(appTags.AppID, db) != AppStatusDoesNotExist
}

func createApp(appTags AppTags, playtime uint32) App {
	return App{AppID: appTags.AppID, Playtime: playtime, Tags: appTags.Tags}
}

func getAppTags(appPlaytime AppPlaytime,
			    returnChan chan App,
			    statusCodeChan chan AppStatus) {
	appTags, err := getAppTagsFromWebsite(appPlaytime.AppID)
	if err != nil {
		fmt.Println(err)
		fmt.Println("retrying...")
		time.Sleep(100 * time.Millisecond)
		getAppTags(appPlaytime, returnChan, statusCodeChan)
	}

	if len(appTags.Tags) <= 0 {
		statusCodeChan <- AppStatus{AppID: appTags.AppID, StatusCode: AppStatusDoesNotExist}
	}

	returnChan <- App{AppID: appPlaytime.AppID, Playtime: appPlaytime.Playtime, Tags: appTags.Tags}
}

func getAppTagsFromWebsite(appID uint32) (AppTags, error) {
	fmt.Printf("Fetching tags from website for: %d\n", appID)

	var appTags AppTags
	html, err := GetGameStorePage(appID)
	if err != nil {
		return appTags, err
	}

	stringTags, err := ExtractTagsFromHTML(html)
	if err != nil {
		return appTags, err
	}

	appTags = AppTags{AppID: appID, Tags: stringTags}

	return appTags, nil
}
