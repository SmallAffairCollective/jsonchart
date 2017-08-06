package main

import (
	"encoding/json"
	"errors"

	"github.com/asaskevich/govalidator"
	"github.com/franela/goreq"
)

func parseUrl(url string) (string, error) {
	isUrl := govalidator.IsURL(url)
	if !isUrl {
		return "", errors.New("error: please enter a valid url")
	}

	return url, nil
}

// return json data object from given url
func fetchJson(url string) map[string]interface{} {

	jsonMap := make(map[string]interface{})

	stringRes, err := fetchUrlData(url)
	check(err)

	err = json.Unmarshal([]byte(stringRes), &jsonMap)
	check(err)

	return jsonMap
}

// fetch data from Url
func fetchUrlData(url string) (string, error) {

	result, err := goreq.Request{Uri: url}.Do()
	check(err)

	stringRes, err := result.Body.ToString()
	if err != nil {
		return "", errors.New("error: unable to request data from url")
	}

	return stringRes, nil
}
