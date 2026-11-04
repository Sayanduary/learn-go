package main

import "fmt"

func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(nums[0])
	fmt.Println(len(nums))
	var vals [4]bool
	vals[2] = true
	fmt.Println(vals)
	nums2 := [3]int{1, 2, 3}
	fmt.Println(nums2)

	for index, val := range nums2 {
		fmt.Println(index, val)
	}
}
