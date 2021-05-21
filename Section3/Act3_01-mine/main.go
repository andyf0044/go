package main

import "fmt"

func main() {
	Costs := map[string]float64{
		"Cake":   0.99,
		"Milk":   2.75,
		"Butter": 0.87,
	}

	SalesTax := map[string]float64{
		"Cake":   7.5,
		"Milk":   1.5,
		"Butter": 2,
	}

	for item, price := range Costs {
		taxrate := SalesTax[item]
		//fmt.Println("Item:", item, "Cost", price, "Tax", taxrate/100, " Sales Tax Total:", price*(taxrate/100))
		fmt.Println("Item:", item, " Sales Tax Total:", price*(taxrate/100))
	}

}
