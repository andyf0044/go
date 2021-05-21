package main

import (
	"fmt"
	"time"
)

func GetConfig() (bool, string, time.Time) {
	return false, "warning", time.Now()
}

func main() {
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)

	Debug2, LogLevel2, startUpTime2 := GetConfig()
	fmt.Println(Debug2, LogLevel2, startUpTime2)
}
