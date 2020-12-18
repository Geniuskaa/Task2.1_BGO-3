package credit

import "math"

func CalculateMonthlyPayment(sumOfCredit float64, percent float64, creditTermInMonths int64) (int64, int64, int64) {
	monthlyPercent := percent / 100 / 12  // Переводим проценты в дробь и ежегодный переводим в месячный
	something := math.Pow((1.0 + monthlyPercent), float64(creditTermInMonths))  // Заменили формулу (1+i)^n
	k := (monthlyPercent * something) / float64(something - 1)
	monthlyPayment := int64(k * sumOfCredit)
	overpayment := int64(float64(monthlyPayment) * float64(creditTermInMonths)) - int64(sumOfCredit)
	totalPayment := int64(float64(monthlyPayment) * float64(creditTermInMonths))

	return monthlyPayment, overpayment, totalPayment
}
