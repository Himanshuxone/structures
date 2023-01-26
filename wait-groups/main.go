package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg = &sync.WaitGroup{}
	fmt.Println("Wait groups for waiting the go routines to be finished")
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		worker(wg, i)
	}
	wg.Wait()
}

func worker(wg *sync.WaitGroup, i int) {
	fmt.Println("worker", i, "waiting for job to finished")
	wg.Done()
}
