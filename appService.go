package main

import (
  "github.com/jinzhu/gorm"
  "encoding/json"
  "regexp"
  "strings"
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

// GetAppTags returns the tags for a given appid
func GetAppTags(appID uint32, db *gorm.DB) ([]Tag, error) {
  if tags := GetAppTagsFromDatabase(appID, db); len(tags) > 0 {
    return tags, nil
  }

  tags, err := getAppTagsFromWebsite(appID)
  if err != nil {
    return nil, err
  }

  InsertTagsIntoDatabase(tags, db)

  return tags, nil
}

func getAppTagsFromWebsite(appID uint32) ([]Tag, error) {
  html, err := GetGameStorePage(appID)
  if err != nil {
    return nil, err
  }

  stringTags, err := extractTagsFromHTML(html)
  if err != nil {
    return nil, err
  }

  tags := make([]Tag, len(stringTags))
  for i, stringTag := range stringTags {
    tags[i] = Tag{Appid: appID, Tag: stringTag}
  }

  return tags, nil
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
