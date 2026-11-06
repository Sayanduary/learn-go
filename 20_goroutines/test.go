package main

import (
	"fmt"
	"sync"
)

func printNum() int {
	sum := 0
	for i := 0; i < 10; i++ {
		fmt.Println("iteration:", i)
		sum += i
	}
	return sum
}

func printLetter() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Printf("%c\n", i)
	}
}

func main() {
	result := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		result <- printNum()
	}()

	go func() {
		defer wg.Done()
		printLetter()
	}()

	go func() {
		wg.Wait()
		close(result)
	}()

	for res := range result {
		fmt.Println(res)
	}
	fmt.Println("finished")
}
