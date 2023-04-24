package solidprinciple

/*
The Liskov Substitution Principle states
that objects of a derived class should be able to replace
objects of the base class without affecting the correctness of the program.
*/

type PaymentProcessor interface {
	ProcessPayment(amount float64) error
}

type CreditCardProcessor struct{}

func (c *CreditCardProcessor) ProcessPayment(amount float64) (err error) {
	// Process credit card payment
	return
}

type PayPalProcessor struct{}

func (p *PayPalProcessor) ProcessPayment(amount float64) (err error) {
	// Process PayPal payment
	return
}

func processPayment(processor PaymentProcessor, amount float64) {
	processor.ProcessPayment(amount)
}

func CallLSP() {
	creditCardProcessor := &CreditCardProcessor{}
	payPalProcessor := &PayPalProcessor{}

	processPayment(creditCardProcessor, 100)
	processPayment(payPalProcessor, 200)
}
