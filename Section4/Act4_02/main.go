package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid no of arguments")
		os.Exit(1)
	}

	users := map[string]string{
		"305": "Sue",
		"204": "Bob",
		"631": "Jake",
		"073": "Tracy",
	}

	// userId, error := strconv.Atoi(os.Args[1])
	// if error != nil {
	// 	fmt.Println("Error with user id", error.Error())
	// } else {
	// 	userName := users[userId]
	// 	if len(userName) == 0 {
	// 		fmt.Println("Unknown user id", userId)
	// 	} else {
	// 		fmt.Println("Hi,", userName)
	// 	}
	// }
	userName := users[os.Args[1]]
	if len(userName) == 0 {
		fmt.Println("Unknown user id", os.Args[1])
	} else {
		fmt.Println("Hi,", userName)
	}
}
