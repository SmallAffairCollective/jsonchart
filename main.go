package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/kataras/iris"
	"github.com/urfave/cli"
)

func main() {
	var delay int
	var iterations int
	var redisHost string
	var url string

	app := cli.NewApp()
	app.Name = "jsonchart"
	app.Usage = "generate charts from JSON endpoints"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "delay",
			Value:       1,
			Usage:       "delay in `SECONDS` between polling URL",
			Destination: &delay,
		},
		cli.IntFlag{
			Name:        "iterations",
			Value:       1,
			Usage:       "number of times to poll URL, -1 runs indefinitely",
			Destination: &iterations,
		},
		cli.StringFlag{
			Name:        "redis",
			Value:       "redis",
			Usage:       "redis `HOST` to connect to",
			Destination: &redisHost,
		},
		cli.StringFlag{
			Name:        "url",
			Usage:       "`URL` to poll from",
			Destination: &url,
		},
	}

	app.Run(os.Args)

	if url != "" {
		url, delay, er := validateArgs(url, delay)
		if er != nil {
			fmt.Println(er.Error())
			return
		}

		go alwaysBeGettin(url, delay, iterations, redisHost)
	} else {
		writeGChartHTML()
		// TODO write JS of existing data in redis
	}

	serve()

}

func serve() {

	app := iris.Default()

	// handle js and html files
	app.StaticWeb("/", "./www")
	app.Run(iris.Addr(":8080"))
}

func alwaysBeGettin(url string, delay int, iterations int, redisHost string) {

	flattenedMatrix := make(map[string]map[string][]float64)
	i := iterations
	counter := 1
	for i != 0 {
		jsonMap := fetchJSON(url)
		metrics := getMetrics(jsonMap)
		conn := connectRedis(redisHost)
		storeMetrics(url, conn, metrics)
		matrix := getStoredMetricMatrix(url, conn)
		flattenedMatrix = flattenMetricMatrix(url, matrix)
		defer conn.Close()
		if iterations == -1 {
			fmt.Println("completed iteration: ", counter, "/ inf")
		} else {
			fmt.Println("completed iteration: ", i, "/", iterations)
		}
		time.Sleep(time.Second * time.Duration(delay))
		writeGChartHTML()
		writeGChartJs(url, delay, counter, flattenedMatrix[url])
		i--
		counter++
	}

}

func validateArgs(url string, delay int) (string, int, error) {

	if delay < 1 {
		return "", -1, errors.New("error: delay must be a number greater than 0")
	}

	// Validate URL
	url, er := parseURL(url)
	if er != nil {
		return "", -1, errors.New(er.Error())
	}

	// Make sure data can be retrieved from URL
	s, er := fetchURLData(url)
	if er != nil {
		return "", -1, errors.New(er.Error())
	}

	// Validate JSON
	var js map[string]interface{}
	if json.Unmarshal([]byte(s), &js) != nil {
		return "", -1, errors.New("error: url does not return valid json")
	}

	// TODO validate redisHost

	return url, delay, nil
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
