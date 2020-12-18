package main

import (
	"fmt"
	"github.com/Geniuskaa/Task2.1_BGO-3/pkg/credit"
)

func main() {
	var sumOfCredit = 1_000_000.00
	var annualPercentOfCredit = 20.0
	var creditTermInMonths int64 = 36

	monthlyPayment, overPayment, totalPayment := credit.CalculateMonthlyPayment(sumOfCredit,annualPercentOfCredit,creditTermInMonths)
	fmt.Println("Ежемесячный платёж составит: ", monthlyPayment, " рублей")
	fmt.Println("Переплата по кредиту составит: ", overPayment, " рублей")
	fmt.Println("Суммарная выплата по кредиту составит: ", totalPayment, " рублей")
}
