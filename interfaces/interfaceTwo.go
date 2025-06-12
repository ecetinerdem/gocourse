package interfaces

import "fmt"

type Mortgage struct {
	creditPaymentTotal float64
	address            string
	rate               float64
}

type Car struct {
	creditPaymentTotal float64
	carInfo            string
	rate               float64
}

type CreditCalculator interface {
	Calculate() float64
}

func (m Mortgage) Calculate() float64 {
	return m.creditPaymentTotal * m.rate / 12
}

func (c Car) Calculate() float64 {
	return c.creditPaymentTotal * c.rate / 12
}

func CalculateMonthlyPayment(credits []CreditCalculator) float64 {
	paymentTotal := 0.0

	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].Calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 1000, address: "Mah"}
	credit2 := Car{rate: 5, creditPaymentTotal: 500, carInfo: "A car"}

	credits := []CreditCalculator{credit1, credit2}
	total := CalculateMonthlyPayment(credits)

	fmt.Println("Total: ", total)
}
