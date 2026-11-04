package main

import (
	"fmt"
	"math/rand"
	"time"
)

// send
func processNum(numChan chan int) {
	for num := range numChan {
		fmt.Println("processing number", num)
		time.Sleep(time.Second * 1)
	}
}

// recieve
func sum(result chan int, num1 int, num2 int) {
	numResult := num1 + num2
	result <- numResult
}

func task(done chan bool) {
	defer func() { done <- true }()
	fmt.Println("processing...")
}

func main() {
	// messageChannel := make(chan string)
	// messageChannel <- "ping"
	// msg := <-messageChannel
	// fmt.Println(msg)
	//
	//
	//
	done := make(chan bool)
	result := make(chan int)
	numChan := make(chan int)

	go task(done)

	<-done
	go sum(result, 4, 5)
	res := <-result
	fmt.Println(res)

	go processNum(numChan)

	for {
		numChan <- rand.Intn(100)
	}

}
