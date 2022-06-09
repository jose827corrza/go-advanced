package main

import "fmt"

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	p.Pay()
}

type CashPayment struct{}

func (CashPayment) Pay() {
	fmt.Println("Paying with cash")
}

type BankPayment struct{}

func (BankPayment) Pay(bankAccount int) { //<- esto da problemas si no se usa el adapter, por que no se usa bien la interface
	fmt.Printf("Paying using bankAccount: %d\n", bankAccount)
}

/*
Aqui empieza la magia del adapter
*/
type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	//bank := &BankPayment{}
	//ProcessPayment(bank) <- asi daria error, por que no usa el adapter
	bpa := &BankPaymentAdapter{
		bankAccount: 1010,
		BankPayment: &BankPayment{},
	}
	ProcessPayment(bpa)
}
