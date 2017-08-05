package main

import (
	"fmt"
)

func printMatrix(flattenedMatrix map[string]map[string][]float64) {
	for item := range flattenedMatrix {
		fmt.Print(item, " ")
		for field := range flattenedMatrix[item] {
			fmt.Print(field, ": ")
			fmt.Print(flattenedMatrix[item][field], " ")
		}
		fmt.Println()
	}
}
