package main

import (
	"io/ioutil"
	"net/http"
)

// GetURL issues a get request to the given url and reads the response
func GetURL(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
