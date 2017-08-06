package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	arg1 := os.Args[1] // url from cmd
	arg2 := os.Args[2] // delay from cmd in secs
	arg3 := os.Args[3] // iterations from cmd

	url, delay, iterations, er := validateArgs(arg1, arg2, arg3)
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	alwaysBeGettin(url, delay, iterations)
}

func alwaysBeGettin(url string, delay int, iterations int) {

	flattenedMatrix := make(map[string]map[string][]float64)
	i := 1
	for i <= iterations {
		jsonMap := fetchJson(url)
		metrics := getMetrics(jsonMap)
		conn := connectRedis("redis")
		storeMetrics(url, conn, metrics)
		matrix := getStoredMetricMatrix(url, conn)
		flattenedMatrix = flattenMetricMatrix(url, matrix)
		defer conn.Close()
		fmt.Println("completed iteration: ", i, "/", iterations)
		time.Sleep(time.Second * time.Duration(delay))
		i++
	}
	writeGChartHtml()
	writeGChartJs(url, delay, iterations, flattenedMatrix[url])
}

func validateArgs(arg1 string, arg2 string, arg3 string) (string, int, int, error) {
	// convert delay and iterations from strings to ints
	delay, er := strconv.Atoi(arg2)
	if er != nil {
		return "", -1, -1, errors.New("error: delay must be a number greater than 0")
	}
	iterations, er := strconv.Atoi(arg3)
	if er != nil {
		return "", -1, -1, errors.New("error: iterations must be a number greater than 0")
	}

	// Make sure delay argument is not less than 1
	if delay < 1 {
		return "", -1, -1, errors.New("error: delay must be greater than 0")
	}

	// Validate URL
	url, er := parseUrl(arg1)
	if er != nil {
		return "", -1, -1, errors.New(er.Error())
	}

	// Make sure data can be retrieved from URL
	s, er := fetchUrlData(url)
	if er != nil {
		return "", -1, -1, errors.New(er.Error())
	}

	// Validate JSON
	var js map[string]interface{}
	if json.Unmarshal([]byte(s), &js) != nil {
		return "", -1, -1, errors.New("error: url does not return valid json")
	}

	return url, delay, iterations, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
