package main

import "fmt"

func CalcSalesTax(ItemCost float64, TaxRate float64) float64 {
	return ItemCost * TaxRate
}

func main() {
	TotalTax := 0.00
	// Cake
	TotalTax += CalcSalesTax(0.99, 0.075)
	// Milk
	TotalTax += CalcSalesTax(2.75, 0.015)
	// Butter
	TotalTax += CalcSalesTax(0.87, 0.02)

	fmt.Println(" Sales Tax Total:", TotalTax)
}
