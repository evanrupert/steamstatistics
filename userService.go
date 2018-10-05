package main

import (
	"encoding/json"
	"errors"
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

	steamID, err := steamIDFromResponse(resp)
	if err != nil {
		return "", err
	}

	return steamID, nil
}

func steamIDFromResponse(resp []byte) (string, error) {
	steamIDResponse := SteamIDResponse{}
	json.Unmarshal(resp, &steamIDResponse)

	if steamIDResponse.Response.Success == 42 {
		return "", errors.New("vanity url not found")
	} else {
		return steamIDResponse.Response.Steamid, nil
	}
}
