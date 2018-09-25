package main

import (
	"encoding/json"
	"net/http"
)

type okResult struct {
	Ok   bool   `json:"ok"`
	Data string `json:"data"`
}

type errResult struct {
	Ok    bool   `json:"ok"`
	Error string `json:"error"`
}

// HelloWorld handler for /
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	sendResponse("Hello, World!", w)
}

func sendResponse(data string, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(okResult{Ok: true, Data: data})
}

func sendError(err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(errResult{Ok: false, Error: err.Error()})
}
