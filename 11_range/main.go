package main

import "fmt"

func main() {
	var nums = make([]int, 0, 2)
	nums = append(nums, 1)
	nums = append(nums, 2)

	nums2 := []int{3, 4, 5}
	var sum int = 0
	for _, val := range nums2 {
		sum = sum + val
	}
	fmt.Println(sum)
	m := map[string]string{"name": "john"}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
