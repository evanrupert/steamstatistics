package main

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/jinzhu/gorm"
)

type userAppsResponse struct {
	Response apps `json:"response"`
}

type apps struct {
	GameCount int           `json:"game_count"`
	Games     []AppPlaytime `json:"games"`
}

// AppPlaytime represents an appid and a user's playtime
type AppPlaytime struct {
	AppID    uint32 `json:"appid"`
	Playtime uint32 `json:"playtime_forever"`
}

// TagPlaytime represents a tag with an associated total playtime
type TagPlaytime struct {
	Tag      string  `json:"tag"`
	Playtime float32 `json:"playtime"`
}

// App represents a whole app object
type App struct {
	AppID    uint32
	Playtime uint32
	Tags     []string
}

// AppTags represents an appid with associated tags
type AppTags struct {
	AppID uint32   `json:"appid"`
	Tags  []string `json:"tags"`
}

// GetTagPlaytimes returns the total playtime for each tag
func GetTagPlaytimes(steamID string) ([]TagPlaytime, error) {
	apps, err := getAllUserAppTags(steamID)
	if err != nil {
		return nil, err
	}

	m := make(map[string]uint32)
	for _, app := range apps {
		for _, tag := range app.Tags {
			if _, ok := m[tag]; ok {
				m[tag] += app.Playtime
			} else {
				m[tag] = app.Playtime
			}
		}
	}

	i := 0
	tagPlaytimes := make([]TagPlaytime, len(m))
	for key, val := range m {
		playtimeHours := minutesToHours(val)
		tagPlaytimes[i] = TagPlaytime{Tag: key, Playtime: playtimeHours}
		i++
	}

	sort.Slice(tagPlaytimes[:], func(i, j int) bool {
		return tagPlaytimes[i].Playtime > tagPlaytimes[j].Playtime
	})

	return tagPlaytimes, nil
}

func minutesToHours(minutes uint32) float32 {
	return float32(minutes) / 60.0
}

func getAllUserAppTags(steamID string) ([]App, error) {
	db, err := OpenConnection()

	if err != nil {
		return nil, err
	}

	apps, err := GetUserApps(steamID)

	if err != nil {
		return nil, err
	}

	appTagsArray := make([]App, len(apps))
	for i, app := range apps {
		appTags, err := GetAppTags(app.AppID, db)
		if err != nil {
			return nil, err
		}

		appTagsArray[i] = App{AppID: appTags.AppID, Playtime: app.Playtime, Tags: appTags.Tags}
	}

	return appTagsArray, nil
}

func GetUserApps(steamID string) ([]AppPlaytime, error) {
	parameters := map[string]string{"steamid": steamID, "format": "json"}
	resp, err := CallMethod("IPlayerService", "GetOwnedGames", 1, parameters)

	if err != nil {
		return nil, err
	}

	apps := appsFromResponse(resp)

	return apps, nil
}

func GetAppTags(appID uint32, db *gorm.DB) (AppTags, error) {
	appTags := GetAppTagsFromDatabase(appID, db)

	if len(appTags.Tags) <= 0 && GetAppStatusCode(appID, db) != AppStatusDoesNotExist {
		var err error
		appTags, err = getAppTagsFromWebsite(appID)

		if len(appTags.Tags) <= 0 {
			InsertAppStatusCode(appTags.AppID, AppStatusDoesNotExist, db)
		}

		if err != nil {
			return appTags, err
		}

		InsertTagsIntoDatabase(appTags, db)
	}

	return appTags, nil
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

func appsFromResponse(resp []byte) []AppPlaytime {
	userAppsResponse := userAppsResponse{}
	json.Unmarshal(resp, &userAppsResponse)

	return userAppsResponse.Response.Games
}
