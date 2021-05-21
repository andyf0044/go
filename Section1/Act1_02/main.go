package main

import "fmt"

func main() {
	a, b := 5, 10
	// call swap here
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
}

func swap(a *int, b *int) {
	// swap the values here
	tmp := *a
	*a = *b
	*b = tmp
	// Quick Go style !
	// *a, *b = *b, *a
}
