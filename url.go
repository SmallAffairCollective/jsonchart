package main

import (
	"encoding/json"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/franela/goreq"
)

func parseURL(url string) (string, error) {
	isURL := govalidator.IsURL(url)
	if !isURL {
		return "", errors.New("error: please enter a valid url")
	}

	return url, nil
}

// return json data object from given url
func fetchJSON(url string) map[string]interface{} {

	jsonMap := make(map[string]interface{})

	stringRes, err := fetchURLData(url)
	check(err)

	err = json.Unmarshal([]byte(stringRes), &jsonMap)
	check(err)

	return jsonMap
}

// fetch data from Url
func fetchURLData(url string) (string, error) {

	result, err := goreq.Request{Uri: url}.Do()
	check(err)

	stringRes, err := result.Body.ToString()
	if err != nil {
		return "", errors.New("error: unable to request data from url")
	}

	return stringRes, nil
}
