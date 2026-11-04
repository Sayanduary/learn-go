package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Sunday:
		fmt.Println("Monday")
	default:
		fmt.Println("Hello")
	}
	whoami := func(i interface{}) {
		switch t := i.(type) {
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		case bool:
			fmt.Println("It is a boolean")
		default:
			fmt.Println("other", t)
		}

	}
	whoami("hello")
}
