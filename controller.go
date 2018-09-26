package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type okResponse struct {
	Ok   bool        `json:"ok"`
	Data interface{} `json:"data"`
}

type errResponse struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// APIRootController handler for api root path: /api
func APIRootController(w http.ResponseWriter, r *http.Request) {
	steamid, err := GetSteamIDFromVanityURL("Eguy45")

	fmt.Println(steamid)

	if err != nil {
		sendError(err, w)
		return
	}

	allAppsTags, err := GetAllUserAppTags(steamid)

	if err != nil {
		sendError(err, w)
		return
	}

	sendResponse(allAppsTags, w)
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
