package main

import (
	"fmt"
	"time"
)

func main() {
	switch dayBorn := time.Sunday; {
	case dayBorn == time.Saturday || dayBorn == time.Sunday:
		fmt.Println("Born on the weekend")
	default:
		fmt.Println("Born on some other day")
	}
}
