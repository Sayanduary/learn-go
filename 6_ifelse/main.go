package main

import "fmt"

func main(){
	var age int = 10;
	if age >= 18 {
		fmt.Println("Person is an adult")
	}else{
		fmt.Println("Not an Adult")
	}

	var role string = "admin"
	var hasPermission bool = true

	if role=="admin" && hasPermission {
		fmt.Println("Logged In")
	}else{
		fmt.Println("Unauthorized")
	}
}
