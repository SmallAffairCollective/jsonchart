package main

import (
	"os"
	"strconv"
	"time"
)

func main() {
	arg1 := os.Args[1] // url from command line
	arg2 := os.Args[2] // delay from cmd in secs
	arg3 := os.Args[3] // iterations from cmd

	// convert delay and iterations from strings to ints
	delay, er := strconv.Atoi(arg2)
	check(er)
	iterations, er := strconv.Atoi(arg3)
	check(er)

	// TODO write regex to make sure arg1 is a valid url

	alwaysBeGettin(arg1, delay, iterations)
}

func alwaysBeGettin(url string, delay int, iterations int) {

	i := 1
	for i <= iterations {
		jsonMap := fetchUrl(url)
		metrics := getMetrics(jsonMap)
		conn := connectRedis("redis")
		storeMetrics(url, conn, metrics)
		matrix := getStoredMetricMatrix(url, conn)
		flattenedMatrix := flattenMetricMatrix(url, matrix)
		printMatrix(flattenedMatrix)
		defer conn.Close()
		time.Sleep(time.Second * time.Duration(delay))
		i++
	}
	writeGChartHtml()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
