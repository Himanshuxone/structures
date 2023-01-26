package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for i := 0; i < 5; i++ {
		go worker(i, jobs, results)
	}

	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs)

	for a := 0; a < 5; a++ {
		value := <-results
		fmt.Println(value)
	}
	close(results)
}

func workerpool() {
	//create worker pool in go?
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
