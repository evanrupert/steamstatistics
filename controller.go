package main

import (
	"encoding/json"
	"net/http"
)

type okResult struct {
	Ok   bool  `json:"ok"`
	Data []App `json:"data"`
}

type errResult struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// APIRoot handler for api root path: /api
func APIRoot(w http.ResponseWriter, r *http.Request) {
	steamID, err := GetSteamIDFromVanityURL("Eguy45")

	if err != nil {
		sendError(err, w)
	}

	userApps, err := GetUserApps(steamID)

	if err != nil {
		sendError(err, w)
	}

	sendResponse(userApps, w)
}

func sendResponse(data []App, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(okResult{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResult{Ok: false, Error: err.Error()})
}
