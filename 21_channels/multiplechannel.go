package main

import "fmt"

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "ping"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("received from chan1:", chan1Val)

		case chan2Val := <-chan2:
			fmt.Println("received from chan2:", chan2Val)
		}
	}
}
