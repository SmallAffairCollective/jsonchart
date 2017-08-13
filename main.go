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

	redisHost := "redis"
	if len(os.Args) > 4 {
		redisHost = os.Args[4] // redis host
	}

	url, delay, iterations, er := validateArgs(arg1, arg2, arg3)
	if er != nil {
		fmt.Println(er.Error())
		return
	}

	alwaysBeGettin(url, delay, iterations, redisHost)
}

func alwaysBeGettin(url string, delay int, iterations int, redisHost string) {

	flattenedMatrix := make(map[string]map[string][]float64)
	i := 1
	for i <= iterations {
		jsonMap := fetchJSON(url)
		metrics := getMetrics(jsonMap)
		conn := connectRedis(redisHost)
		storeMetrics(url, conn, metrics)
		matrix := getStoredMetricMatrix(url, conn)
		flattenedMatrix = flattenMetricMatrix(url, matrix)
		defer conn.Close()
		fmt.Println("completed iteration: ", i, "/", iterations)
		time.Sleep(time.Second * time.Duration(delay))
		i++
	}
	writeGChartHTML()
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
	url, er := parseURL(arg1)
	if er != nil {
		return "", -1, -1, errors.New(er.Error())
	}

	// Make sure data can be retrieved from URL
	s, er := fetchURLData(url)
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
