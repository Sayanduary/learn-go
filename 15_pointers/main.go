package main

import "fmt"

func changeNum(num *int) {
	*num = 5
	fmt.Println("IN CHANGE :", *num)
}

func main() {
	num := 1
	changeNum(&num)
	fmt.Println(num)
}
