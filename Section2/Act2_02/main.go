package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	favouriteWord := ""
	favouriteCount := 0

	for key, count := range words {
		if count > favouriteCount {
			favouriteCount = count
			favouriteWord = key
		}
	}
	fmt.Println("Most popular word:", favouriteWord)
	fmt.Println("With a count of:", favouriteCount)
}
