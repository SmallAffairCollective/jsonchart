package main

import (
	"bufio"
	"fmt"
	"os"
)

func queryUser(fields []string) []string {
	var requestedFields []string
	reader := bufio.NewReader(os.Stdin)

	for _, f := range fields {
		fmt.Printf("Include field '%s'? [Y/n] ", f)
		answer, _ := reader.ReadString('\n')
		fmt.Println("Yous said: ", answer)
		if answer == "Y\n" || answer == "y\n" || answer == "\n" {
			requestedFields = append(requestedFields, f)
		}
	}
	return requestedFields
}
