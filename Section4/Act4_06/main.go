package main

import (
	"fmt"
)

func typeChecker(v interface{}) (string, error) {
	switch t := v.(type) {
	case string:
		return fmt.Sprint("string:", t), nil
	case int, int32, int64:
		return fmt.Sprint("int:", t), nil
	case float32, float64:
		return fmt.Sprint("float:", t), nil
	case bool:
		return fmt.Sprint("bool:", t), nil
	default:
		return "unknown", nil
	}
	//	return "", errors.New("Error when checking type")
}

func main() {

	// featureVector := []interface{}{
	// 	[]int{1, 2},
	// 	[]int8{10, 12},
	// 	[]int16{1, 3},
	// 	[]int32{200, 1000},
	// 	[]int64{100, 2000},
	// 	[]float32{2.5, -3.0},
	// 	[]float64{1.2, 2.2},
	// 	[]string{"a", "b"},
	// 	[]bool{false, true},
	// 	[]uint16{5, 12},
	// 	[]uint32{99, 999},
	// }

	// for i := 0; i < len(featureVector); i++ {
	// 	typeStr, err := typeChecker(featureVector[i])
	// 	if err == nil {
	// 		fmt.Println("Element:", i, " type:", typeStr)
	// 	} else {
	// 		fmt.Println("Element:", i, " error:", err)
	// 	}
	// }

	var i1 int = 1
	typeStr, err := typeChecker(i1)
	if err == nil {
		fmt.Println("Element: 1 type:", typeStr)
	} else {
		fmt.Println("Element: 1 error:", err)
	}

	var i2 int8 = 64
	typeStr, err = typeChecker(i2)
	if err == nil {
		fmt.Println("Element: 2 type:", typeStr)
	} else {
		fmt.Println("Element: 2 error:", err)
	}

	var i3 int32 = 101
	typeStr, err = typeChecker(i3)
	if err == nil {
		fmt.Println("Element: 3 type:", typeStr)
	} else {
		fmt.Println("Element: 3 error:", err)
	}

	var i4 int64 = -1999
	typeStr, err = typeChecker(i4)
	if err == nil {
		fmt.Println("Element: 4 type:", typeStr)
	} else {
		fmt.Println("Element: 4 error:", err)
	}

	var i5 float32 = 10
	typeStr, err = typeChecker(i5)
	if err == nil {
		fmt.Println("Element: 5 type:", typeStr)
	} else {
		fmt.Println("Element: 5 error:", err)
	}

	var i6 float64 = 10
	typeStr, err = typeChecker(i6)
	if err == nil {
		fmt.Println("Element: 6 type:", typeStr)
	} else {
		fmt.Println("Element: 6 error:", err)
	}

	var i7 bool = true
	typeStr, err = typeChecker(i7)
	if err == nil {
		fmt.Println("Element: 7 type:", typeStr)
	} else {
		fmt.Println("Element: 7 error:", err)
	}

	i8 := struct {
		x int
		y int
	}{
		x: 2,
		y: 8,
	}
	typeStr, err = typeChecker(i8)
	if err == nil {
		fmt.Println("Element: 8 type:", typeStr)
	} else {
		fmt.Println("Element: 4 error:", err)
	}

	i9 := "Hello!"
	typeStr, err = typeChecker(i9)
	if err == nil {
		fmt.Println("Element: 9 type:", typeStr)
	} else {
		fmt.Println("Element: 9 error:", err)
	}

}
