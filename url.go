package main

import (
	"encoding/json"

	"github.com/franela/goreq"
)

func parseUrl(url string) string {
	// error check to make sure it's a valid url

	return url
}

// return data object from given url
func fetchUrl(url string) (map[string]interface{}, string) {

	result, err := goreq.Request{Uri: url}.Do()
	check(err)

	jsonMap := make(map[string]interface{})

	stringRes, err := result.Body.ToString()
	check(err)

	err = json.Unmarshal([]byte(stringRes), &jsonMap)
	check(err)

	return jsonMap, stringRes
}
