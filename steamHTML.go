package main

import (
	"fmt"
	"regexp"
	"strings"
)

const baseStoreURL = "https://store.steampowered.com"

// GetGameStorePage returns the html of the store page for the given appid
func GetGameStorePage(appID uint32) ([]byte, error) {
	url := fmt.Sprintf("%s/app/%d", baseStoreURL, appID)

	resp, err := GetURLWithCookie(url)

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func ExtractTagsFromHTML(html []byte) ([]string, error) {
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
