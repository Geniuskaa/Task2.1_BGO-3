package credit_test

import (
	"fmt"
	"github.com/Geniuskaa/Task2.1_BGO-3/pkg/credit"
)

func ExampleCalculateMonthlyPayment() {
	fmt.Println(credit.CalculateMonthlyPayment(1_000_000, 20, 36))
	fmt.Println(credit.CalculateMonthlyPayment(100_000, 20, 36))
	// Output:
	// 37163 337868 1337868
	// 3716 33776 133776
}
