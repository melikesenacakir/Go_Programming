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
	calculate() float64
}

func (m Mortgage) calculate() float64 {
	return m.creditPaymentTotal * m.rate / 1200
}

func (c Car) calculate() float64 {
	return c.creditPaymentTotal * c.rate / 1200
}

func MonthlyPaymentCalculate(credits []CreditCalculator) float64 {
	paymentTotal := 0.0
	for i := 0; i < len(credits); i++ {
		paymentTotal += credits[i].calculate()
	}
	return paymentTotal
}

func Demo2() {
	credit1 := Mortgage{rate: 10, creditPaymentTotal: 100000, address: "Ankara"}
	credit2 := Mortgage{rate: 12, creditPaymentTotal: 5000, address: "Kocaeli"}
	credit3 := Car{rate: 15, creditPaymentTotal: 6000, carInfo: "Polo"}
	credits := []CreditCalculator{credit1, credit2, credit3}
	total := MonthlyPaymentCalculate(credits)
	fmt.Println("Toplam Ödeme", total)
}
