package main

import (
	"encoding/json"
)

// SteamIDResponse represents the response of a call to ResolveVanityURL
type SteamIDResponse struct {
	Response SteamID `json:"response"`
}

// SteamID represents the steamid and a status code
type SteamID struct {
	Steamid string `json:"steamid"`
	Success int    `json:"success"`
}

// GetSteamIDFromVanityURL finds the steamID given a vanityURL
func GetSteamIDFromVanityURL(vanityURL string) (string, error) {
	parameters := map[string]string{"vanityurl": vanityURL}
	resp, err := CallMethod("ISteamUser", "ResolveVanityURL", 1, parameters)

	if err != nil {
		return "", err
	}

	steamID := steamIDFromResponse(resp)

	return steamID, nil
}

func steamIDFromResponse(resp []byte) string {
	steamIDResponse := SteamIDResponse{}
	json.Unmarshal(resp, &steamIDResponse)

	return steamIDResponse.Response.Steamid
}
