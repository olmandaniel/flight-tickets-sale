package services

import "fmt"

type Payable interface {
	Pay() error
}

type CreditCardDriver struct {
}

func (CreditCardDriver) Pay() error {
	fmt.Println("Pay with credit card")
	return nil
}

type GooglePayDriver struct {
}

func (GooglePayDriver) Pay() error {
	fmt.Println("Pay with google pay")
	return nil
}

type PayPalDriver struct {
}

func (PayPalDriver) Pay() error {
	fmt.Println("Pay with paypal")
	return nil
}

type PaymentProcessorFactory struct {
}

func (PaymentProcessorFactory) GetPaymentDriver(name string) Payable {
	switch name {
	case "CREDIT_CARD":
		return &CreditCardDriver{}
	case "GOOGLE_PAY":
		return &GooglePayDriver{}
	case "PAYPAL":
		return &PayPalDriver{}
	default:
		return nil

	}
}
