package main

import "fmt"

// enumerated types

type OrderStatus string

const (
	RECIEVED = "recieved"
	CONFIRMED="confirmed"
	PREPARED="prepared"
	DELIVERED="delicvered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Order Status to ",status);
}


func main(){
changeOrderStatus(PREPARED)
}