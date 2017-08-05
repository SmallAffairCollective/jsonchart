package main

import (
	"os"
)

func main() {
	url := os.Args[1] // url from command line

	fetchUrl(url)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
