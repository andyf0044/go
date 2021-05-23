package main

import (
	"fmt"
	"os"
	"strconv"
)

func getQAResults() []string {
	data := []string{
		"Good",
		"Good",
		"Bad",
		"Good",
		"Good",
	}
	return data
}

func removeElement(data []string, index int) []string {

	modified := make([]string, 0, len(data)-1)
	for i := 0; i < len(data); i++ {
		if i != index {
			modified = append(modified, data[i])
		}
	}
	return modified
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Data index not passed")
		os.Exit(1)
	}
	badDataIndex, err := strconv.Atoi(os.Args[1])
	results := getQAResults()
	if err != nil {
		fmt.Println("Data index not recognised.", err)
		os.Exit(2)
	}

	if badDataIndex < 0 || badDataIndex >= len(results) {
		fmt.Println("Passed index is not valid")
		os.Exit(3)
	}
	fmt.Println(results)
	results = removeElement(results, badDataIndex)
	fmt.Println(results)
}
