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
	state := storeMetrics(url, "redis", metrics)
	fmt.Println(state)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
