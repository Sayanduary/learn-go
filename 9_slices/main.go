package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums == nil)

	var nums2 = make([]int, 0, 5)
	nums2 = append(nums2, 1)
	nums2 = append(nums2, 2)
	nums2 = append(nums2, 4)
	nums3 := nums2
	fmt.Println(len(nums2))
	fmt.Println(nums3)
	fmt.Println(nums2[0:2])
}
