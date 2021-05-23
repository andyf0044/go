package main

import (
	"fmt"
	"math"
)

func convert() string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14

	// convert from smaller int to larger int type - always safe
	m := fmt.Sprintf("int8 = %v > in64 = %v\n", i8, int64(i8))
	// convert from an int 1 above int8 max size - cause overflow
	m += fmt.Sprintf("int = %v > in8  = %v\n", i, int8(i))
	// convert int8 to float64-doesn't cause overflow
	m += fmt.Sprintf("int8 = %v > float64 = %v\n", i8, float64(i8))
	// convert float into an int, whole no remains, decimal data is lost
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))
	return m
}

func main() {
	fmt.Print(convert())
}
