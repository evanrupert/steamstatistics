package main

import (
	"encoding/json"
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
	db, err := OpenConnection()

	if err != nil {
		sendError(err, w)
	}

	tags, err := GetAppTags(400, db)

	if err != nil {
		sendError(err, w)
	}

	sendResponse(tags, w)
}

// GetTagsController handler for path: /api/tags
func GetTagsController(w http.ResponseWriter, r *http.Request) {
	tags := []string{"Tag1", "Tag2", "Tag3"}

	sendResponse(tags, w)
}

func sendResponse(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(okResponse{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResponse{Ok: false, Error: err.Error()})
}
