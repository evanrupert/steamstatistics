package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/jinzhu/gorm"
)

type userAppsResponse struct {
	Response apps `json:"response"`
}

type apps struct {
	GameCount int   `json:"game_count"`
	Games     []App `json:"games"`
}

// App represents an appid and a user's playtime
type App struct {
	AppID    uint32 `json:"appid"`
	Playtime uint32 `json:"playtime_forever"`
}

// AppTags represents an appid with associated tags
type AppTags struct {
	AppID uint32   `json:"appid"`
	Tags  []string `json:"tags"`
}

// GetAllUserAppTags returns all apptags for a user
func GetAllUserAppTags(steamID string) ([]AppTags, error) {
	db, err := OpenConnection()

	if err != nil {
		return nil, err
	}

  apps, err := getUserApps(steamID)

	if err != nil {
    return nil, err
	}

	appTagsArray := make([]AppTags, len(apps))
	for i, app := range apps {
		appTags, err := getAppTags(app.AppID, db)
		if err != nil {
			return nil, err
		}

		appTagsArray[i] = appTags
	}

	return appTagsArray, nil
}

func getUserApps(steamID string) ([]App, error) {
  fmt.Println(steamID)
	parameters := map[string]string{"steamid": steamID, "format": "json"}
	resp, err := CallMethod("IPlayerService", "GetOwnedGames", 1, parameters)

	if err != nil {
		return nil, err
	}

	apps := appsFromResponse(resp)

	return apps, nil
}

// GetAppTags returns the tags for a given appid
func getAppTags(appID uint32, db *gorm.DB) (AppTags, error) {
	var appTags AppTags

	if appTags = GetAppTagsFromDatabase(appID, db); len(appTags.Tags) > 0 {
		fmt.Printf("Getting tags from database for: %d\n", appID)
	} else {
		appTags, err := getAppTagsFromWebsite(appID)

		if err != nil {
			return appTags, err
		}

		fmt.Printf("Getting tags from website for: %d\n", appID)

		InsertTagsIntoDatabase(appTags, db)
	}

	return appTags, nil
}

func getAppTagsFromWebsite(appID uint32) (AppTags, error) {
	var appTags AppTags
	html, err := GetGameStorePage(appID)
	if err != nil {
		return appTags, err
	}

	stringTags, err := extractTagsFromHTML(html)
	if err != nil {
		return appTags, err
	}

	appTags = AppTags{AppID: appID, Tags: stringTags}

	return appTags, nil
}

func extractTagsFromHTML(html []byte) ([]string, error) {
	regex, err := regexp.Compile(`class="app_tag"[^>]*>([^<]*)`)

	if err != nil {
		return nil, err
	}

	matches := regex.FindAllStringSubmatch(string(html), -1)

	tags := make([]string, len(matches))

	for i, match := range matches {
		tags[i] = strings.TrimSpace(match[1])
	}

	return tags, nil
}

func appsFromResponse(resp []byte) []App {
	userAppsResponse := userAppsResponse{}
	json.Unmarshal(resp, &userAppsResponse)

	return userAppsResponse.Response.Games
}
