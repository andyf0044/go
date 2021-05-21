package main

import "fmt"

const GoodCreditThres = 450

func GetInterestRate(CreditScore int) int {
	if CreditScore < GoodCreditThres {
		return 20
	} else {
		return 15
	}
}

func CheckLoan(CreditScore int, Loan float64, MonthlySalary float64, DurationMonths int) (float64, int, float64, bool) {
	var MonthlyPayment float64 = 0
	var Rate int = 0
	var TotalCost float64 = 0
	var Approved bool = false

	if CreditScore < 0 || Loan < 0 || MonthlySalary < 0 || DurationMonths <= 0 || DurationMonths%12 != 0 {
		return MonthlyPayment, Rate, TotalCost, Approved
	}

	Rate = GetInterestRate(CreditScore)
	TotalInterest := Loan * (float64(Rate) / 100.00000000)
	TotalCost = Loan + TotalInterest
	MonthlyPayment = TotalCost / float64(DurationMonths)

	if CreditScore < GoodCreditThres {
		if MonthlyPayment <= (MonthlySalary * 0.10000000) {
			Approved = true
		}
	} else {
		if MonthlyPayment <= (MonthlySalary * 0.20000000) {
			Approved = true
		}
	}

	return MonthlyPayment, Rate, TotalCost, Approved
}

func main() {
	Name := "John Smith"
	CreditScore := 350
	MonthlyIncome := 1000.00
	LoanVal := 1000.00
	Duration := 12

	var MonthlyPayment float64 = 0.00
	var IntRate int = 0
	var TotalRepayment float64 = 0.00
	var LoanApproved bool = false

	MonthlyPayment, IntRate, TotalRepayment, LoanApproved = CheckLoan(CreditScore, LoanVal, MonthlyIncome, Duration)

	fmt.Println("Applicant", Name)
	fmt.Println("---------------------------------")
	fmt.Println("Credit Score:", CreditScore)
	fmt.Println("Income:", MonthlyIncome)
	fmt.Println("Loan Amount:", LoanVal)
	fmt.Println("Loan Term:", Duration)
	fmt.Println("Monthly Payment:", MonthlyPayment)
	fmt.Println("Rate:", IntRate)
	fmt.Println("Total Cost:", TotalRepayment)
	fmt.Println("Approved:", LoanApproved)

}
