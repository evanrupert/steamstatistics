package main

import (
	"encoding/json"
	"math"
	"sort"
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
	Playtime float64 `json:"playtime"`
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
	apps, err := GetAllUserAppTags(steamID)
	if err != nil {
		return nil, err
	}

	tagPlaytimes := convertTagPlaytimeMapToArray(calculateHoursPerTag(apps))

	sort.Slice(tagPlaytimes[:], func(i, j int) bool {
		return tagPlaytimes[i].Playtime > tagPlaytimes[j].Playtime
	})

	return tagPlaytimes, nil
}

func calculateHoursPerTag(apps []App) map[string]uint32 {
	tagPlaytimeMap := make(map[string]uint32)

	for _, app := range apps {
		for _, tag := range app.Tags {
			if tagExistsInMap(tagPlaytimeMap, tag) {
				tagPlaytimeMap[tag] += app.Playtime
			} else {
				tagPlaytimeMap[tag] = app.Playtime
			}
		}
	}

	return tagPlaytimeMap
}

func convertTagPlaytimeMapToArray(tagPlaytimeMap map[string]uint32) []TagPlaytime {
	index := 0
	tagPlaytimes := make([]TagPlaytime, len(tagPlaytimeMap))
	for tag, playtime := range tagPlaytimeMap {
		playtimeHours := minutesToHours(playtime)
		tagPlaytimes[index] = TagPlaytime{Tag: tag, Playtime: playtimeHours}
		index++
	}

	return tagPlaytimes
}

func tagExistsInMap(tagPlaytimeMap map[string]uint32, tag string) bool {
	_, exists := tagPlaytimeMap[tag]
	return exists
}

func minutesToHours(minutes uint32) float64 {
	hours := float64(minutes) / 60.0
	return roundToHundreds(hours)
}

func roundToHundreds(number float64) float64 {
	unit := 100.0
	return math.Round(number * unit) / unit
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

func appsFromResponse(resp []byte) []AppPlaytime {
	userAppsResponse := userAppsResponse{}
	json.Unmarshal(resp, &userAppsResponse)

	return userAppsResponse.Response.Games
}
