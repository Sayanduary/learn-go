package main

import (
	"fmt"
	"maps"
)

func main() {
	//maps -> Hash , object , dict
	m := make(map[string]int)
	m["name"] = 1
	m["area"] = 2
	clear(m)
	fmt.Println(m["price"], m["qty"])
	n := map[string]int{"price": 40, "qty": 10}
	fmt.Println(n["price"])
	// check wheather the element present in the map
	value, isPresent := n["price"]
	if isPresent {
		fmt.Println("all okay", "Value: ", value)
	} else {
		fmt.Println("Not Okay")
	}
	fmt.Println(maps.Equal(m, n))
}
