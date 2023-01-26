package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)
	done := make(chan int)
	go routines(done)
	go func() {
		messages <- "ping"
	}()
	msg := <-messages
	fmt.Println(msg)
	<-done
}

func routines(done chan int) {
	fmt.Println("Started the message")
	time.Sleep(time.Second)
	fmt.Println("Ended the message")
	done <- 1
}
