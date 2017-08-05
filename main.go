package main

import (
	"fmt"
	"os"
)

func main() {
	url := os.Args[1] // url from command line

	jsonMap := fetchUrl(url)
	metrics := getMetrics(jsonMap)
	fmt.Println(metrics)
	conn := connectRedis("redis")
	state := storeMetrics(url, conn, metrics)
	fmt.Println(state)
	matrix := getStoredMetricMatrix(url, conn)
	fmt.Println(matrix)
	flattenedMatrix := flattenMetricMatrix(url, matrix)
	printMatrix(flattenedMatrix)
	defer conn.Close()
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
