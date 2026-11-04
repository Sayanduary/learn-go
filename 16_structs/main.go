// Custom Data Structures
package main

import (
	"fmt"
	"time"
)

// order struct
type Customer struct{
	name string
	phone string
}


type Order struct{
	id string
	amount float32
	status string
	created_at time.Time // nano seconds precissions
	Customer // struct embedding
}


// constructor
func newOrder(id string, amount float32,status string )*Order{
	// initial setup here..
myOrder:=Order{
id:id,
amount: amount,
status: status,
}
return &myOrder
}

func (O *Order)changeStatus(status string)  {
O.status=status
}

func (O Order) get_amount()float32{
return O.amount
}

func main() {

	language:=struct{
		name string
		isGood bool
	}{"golang",true}

	fmt.Println(language)

		order:=Order{
			id: "1",
			amount: 120.45,
			status: "dispatch",
			Customer: Customer{
				name: "sayan",
				phone: "3423535353",
			},
		}
		order2:=Order{
			id:"2",
			amount:23.45 ,
			status: "Delivered",
			created_at: time.Now(),
		}
		order2.changeStatus("Shipped")

		myOrder:=newOrder("1",120.21,"Delivered"); // build with constructor
		fmt.Println(myOrder)
		fmt.Println(order.id)
		fmt.Println(order2)
		fmt.Println(order2.get_amount())
		fmt.Println(order.Customer.name)
}
