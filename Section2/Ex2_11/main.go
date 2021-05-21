package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for {
		r := rand.Intn(8)
		if r%3 == 0 {
			fmt.Println("Skip", r)
			continue
		} else if r%2 == 0 {
			fmt.Println("Stop", r)
			break
		}
		fmt.Println(r)
	}
}
