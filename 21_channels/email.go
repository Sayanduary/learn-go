// buffer channel
package main

import (
	"fmt"
	"time"
)

func emailSender(emailChan chan string, done chan bool) {

	defer func() { done <- true }() // defer means execute this function when emailSender() is about to finish.

	for email := range emailChan {
		fmt.Println("Sending Email to", email)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	email_chan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(email_chan, done)
	for i := 0; i < 5; i++ {
		email_chan <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending")
	close(email_chan)

	// 	email_chan <- "1@example.com"
	// 	email_chan <- "2@example.com"
	//
	// 	fmt.Println(<-email_chan)
	// 	fmt.Println(<-email_chan)

	<-done
}
