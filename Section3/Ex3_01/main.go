package main

import (
	"fmt"
	"unicode"
)

func passwordChecker(pw string) bool {
	PwRune := []rune(pw)
	if len(PwRune) < 8 {
		return false
	}
	if len(PwRune) > 15 {
		return false
	}
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	for _, v := range PwRune {
		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}
	}
	return hasUpper && hasLower && hasNumber && hasSymbol
}

func main() {
	// if passwordChecker("") {
	// 	fmt.Println("password good")
	// } else {
	// 	fmt.Println("password bad")
	// }
	// if passwordChecker("This!I5A") {
	// 	fmt.Println("password good")
	// } else {
	// 	fmt.Println("password bad")
	// }
	passwds := []string{
		"",
		"Ratatoille",
		"Ratato1ll3",
		"R@t@t01ll3",
		"Ratatoille!",
		"Ratatoille;",
		"Ratat0ille;",
	}

	for _, pw := range passwds {
		if passwordChecker(pw) {
			fmt.Println("password", pw, "good")
		} else {
			fmt.Println("password", pw, "bad")
		}
	}
}
