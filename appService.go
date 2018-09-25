package main

import (
	"encoding/json"
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

// GetUserApps returns all owned games given a user's steamID
func GetUserApps(steamID string) ([]App, error) {
	parameters := map[string]string{"steamid": steamID, "format": "json"}
	resp, err := CallMethod("IPlayerService", "GetOwnedGames", 1, parameters)

	if err != nil {
		return nil, err
	}

	apps := appsFromResponse(resp)

	return apps, nil
}

func appsFromResponse(resp []byte) []App {
	userAppsResponse := userAppsResponse{}
	json.Unmarshal(resp, &userAppsResponse)

	return userAppsResponse.Response.Games
}
