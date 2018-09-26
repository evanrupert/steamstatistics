package main

import (
	"fmt"
)

const baseStoreURL = "https://store.steampowered.com"

// GetGameStorePage returns the html of the store page for the given appid
func GetGameStorePage(appID uint32) ([]byte, error) {
	url := fmt.Sprintf("%s/app/%d", baseStoreURL, appID)

	resp, err := GetURL(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
