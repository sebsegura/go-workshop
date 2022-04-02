package main

import (
	"errors"
	"fmt"
)

const (
	Credit = 1
	Debit  = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

// factory
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Credit:
		return &CreditPM{}, nil
	case Debit:
		return &DebitPM{}, nil
	default:
		return nil, errors.New("Payment method not recognized")
	}
}

type CreditPM struct{}
type DebitPM struct{}

func (c *CreditPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using credit card\n", amount)
}

func (d *DebitPM) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}

func main() {
	cardPayment, err := GetPaymentMethod(Credit)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(cardPayment.Pay(1.5))

	debitPayment, err := GetPaymentMethod(Debit)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(debitPayment.Pay(2.5))
}
