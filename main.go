package main

import (
	"fmt"
	"os"
)

func main() {
	url := os.Args[1] // url from command line

	jsonMap, stringRes := fetchUrl(url)

	// testing, not working yet
	path := make(map[string]string)
	for k := range jsonMap {
		path = goDeeper(stringRes, path, k)
	}
	fmt.Println("path:", path)

	metrics := getMetrics(jsonMap)
	fmt.Println(metrics)
	conn := connectRedis("redis")
	state := storeMetrics(url, conn, metrics)
	fmt.Println(state)
	matrix := getStoredMetricMatrix(url, conn)
	fmt.Println(matrix)
	defer conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
