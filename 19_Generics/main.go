package main

import "fmt"


func printSlice[T any](items[]T){
	for _,val:=range items{
		fmt.Println(val)
	}
}

func main(){
printSlice([]string{"1","2","3"})
}