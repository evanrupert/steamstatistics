package main

import (
  "encoding/json"
  "net/http"
)

type appsResponse struct {
  Ok   bool  `json:"ok"`
  Data []App `json:"data"`
}

type tagsResponse struct {
  Ok bool `json:"ok"`
  Data []string `json:"data"`
}

type errResult struct {
  Ok    bool   `json:"ok"`
  Error string `json:"error"`
}

// APIRootController handler for api root path: /api
func APIRootController(w http.ResponseWriter, r *http.Request) {
  steamID, err := GetSteamIDFromVanityURL("Eguy45")

  if err != nil {
    sendError(err, w)
  }

  userApps, err := GetUserApps(steamID)

  if err != nil {
    sendError(err, w)
  }

  sendAllApps(userApps, w)
}

// GetTagsController handler for path: /api/tags
func GetTagsController(w http.ResponseWriter, r *http.Request) {
  tags := []string{"Tag1", "Tag2", "Tag3"}

  sendTags(tags, w)
}

func sendAllApps(data []App, w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(appsResponse{Ok: true, Data: data})
}

func sendTags(data []string, w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(tagsResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
  w.Header().Set("Content-Type", "application/json")

  json.NewEncoder(w).Encode(errResult{Ok: false, Error: err.Error()})
}
