package main

import "fmt"

func getMondayWeek() []string {
	days := []string{
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
		"Sunday",
	}
	return days
}

func shiftWeekToSunday(days []string) []string {
	newWeek := []string{days[len(days)-1]}
	remainingDays := days[:len(days)-1]
	newWeek = append(newWeek, remainingDays...)
	return newWeek
}

func main() {
	weekdays := getMondayWeek()
	fmt.Println("Start:", weekdays)
	weekdays = shiftWeekToSunday(weekdays)
	fmt.Println("Finish", weekdays)
}
