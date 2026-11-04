package main

import "fmt"


type IPayment interface{
	pay(amount float32)
	refund(amount float32,accoountNo string)
}

type Payment struct{
gateway IPayment
}

type Razorpay struct{

}

func (R Razorpay)pay(amount float32){
	// logic for make payment
	fmt.Println("Payed with razorpay",amount)
}

// type Stripe struct{

// }

// func(S *Stripe) pay(amount float32){
// 	fmt.Println("Payed with stripe",amount)
// }

type FakePayment struct{}

func (F FakePayment)pay(amount float32)  {
fmt.Println("Payed with fake payment gateway",amount)
}

type Paypal struct{}

func (P Paypal)pay(amount float32)  {
	fmt.Println("Payed with Paypal payment gateway",amount)
}
func (P Paypal)refund(amount float32,accountNo string)  {
	fmt.Println("Refunded with Paypal payment gateway",amount,accountNo)
}

func main(){
	PayementGateway:=Paypal{}
	newPayment:=Payment{
		gateway: PayementGateway,
	}
	newPayment.gateway.pay(100)
	newPayment.gateway.refund(200,"q42443345")
}
