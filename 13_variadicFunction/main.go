package main

import "fmt"

func sum(nums ...int) int {
	total := 0
	for _, i := range nums {
		total = total + i
	}
	return total
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(sum(nums...))
}
