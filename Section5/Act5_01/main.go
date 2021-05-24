package main

type Weekday int

const (
	Monday Weekday = iota // Index = 0
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

type Employee struct {
	Id        int
	FirstName string
	LastName  string
}

type developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}

func main() {

}
