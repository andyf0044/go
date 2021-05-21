package main

import "fmt"

var defaultOffset = 10

func main() {
	// offset := 5
	offset := defaultOffset
	fmt.Println(offset)
	offset = offset + defaultOffset
	fmt.Println(offset)
}
