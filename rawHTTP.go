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

// GetURLWithCookie issues a get request to the given url with the birthtime cookie and reads the response
func GetURLWithCookie(url string) ([]byte, error) {
	cookie := http.Cookie{Name: "birthtime", Value: "568022401"}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.AddCookie(&cookie)

	var client = &http.Client{}

	resp, err := client.Do(req)

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
