package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func getLanguages() (string, string, bool) {
	return "go", "js", false
}

func processIt(fn func(a int) int) {
	fn(1)
}

func process() func(a int) int {
	return func(a int) int {
		return a
	}
}
func main() {

	fn := func(a int) int {
		return 2
	}

	fn2 := process()

	processIt(fn)

	// fmt.Println(add(10, 12))
	fmt.Println(fn2(10))
}
