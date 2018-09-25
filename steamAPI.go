package main

import (
	"fmt"
	"os"
)

const baseURL = "http://api.steampowered.com"

// CallMethod calls the given steam api method
func CallMethod(steamInterface string,
	method string,
	version int,
	additionalParameters map[string]string) ([]byte, error) {
	additionalParameters["key"] = os.Getenv("STEAM_API_KEY")

	url := buildRequestURL(steamInterface, method, version, additionalParameters)

	resp, err := GetURL(url)

	if err != nil {
    return nil, err
	}

  return resp, nil
}

func buildRequestURL(steamInterface string,
	method string,
	version int,
	additionalParameters map[string]string) string {
	url := fmt.Sprintf("%s/%s/%s/v00%d/",
		baseURL,
		steamInterface,
		method,
		version)

	url = fmt.Sprintf("%s?", url)

	for k, v := range additionalParameters {
		url = fmt.Sprintf("%s%s=%s&", url, k, v)
	}

	return url
}
