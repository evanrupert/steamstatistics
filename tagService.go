package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
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

	appTagsChan := make(chan AppTags)
	stopDatabaseChan := make(chan bool)

	go DatabaseWriter(appTagsChan, stopDatabaseChan, db)

	i := 0

	for _, app := range apps {
		err := getAppTags(app.AppID, appTagsChan, db)
		if err != nil {
			return nil, err
		}
		i++

		//appTagsChan <- appTags
		//
		//appTagsArray[i] = App{AppID: appTags.AppID, Playtime: app.Playtime, Tags: appTags.Tags}
	}

	// TODO: Finish implementing concurrency for the getAppTags process
	// NOTE: this will have to be redesign because the current algo depends on
	// having access to the app on line 28 but of course there is not order if it is performed concurrently
	for i > 0 {
		appTags := <- appTagsChan
		appTagsArray = append(appTagsArray, /* Insert App here */)
	}

	return appTagsArray, nil
}

func getAppTags(appID uint32, c chan AppTags, db *gorm.DB) error {
	appTags := GetAppTagsFromDatabase(appID, db)

	if len(appTags.Tags) <= 0 && GetAppStatusCode(appID, db) != AppStatusDoesNotExist {
		var err error
		appTags, err = getAppTagsFromWebsite(appID)

		if len(appTags.Tags) <= 0 {
			InsertAppStatusCode(appTags.AppID, AppStatusDoesNotExist, db)
		}

		if err != nil {
			return err
		}
	}

	c <- appTags

	return nil
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
