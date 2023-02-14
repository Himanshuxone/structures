package main

import "fmt"

type stack []int

func main() {
	var stacks stack = make([]int, 0)
	data := []int{1, 2, 3, 4, 5}
	for i := len(data) - 1; i >= 0; i-- {
		stacks.Push(data[i])
	}
	fmt.Printf("Stacks => %+v\n", stacks)
	for _ = range stacks {
		stacks.Pop()
	}
	fmt.Printf("Stacks => %+v\n", stacks)
	for i := len(data) - 1; i >= 0; i-- {
		stacks.Enqueue(data[i])
	}
	fmt.Printf("Stacks => %+v\n", stacks)
	stacks.Dequeue()
	fmt.Printf("Stacks => %+v\n", stacks)
	stacks.Dequeue()
	fmt.Printf("Stacks => %+v\n", stacks)
	stacks.Dequeue()
	fmt.Printf("Stacks => %+v\n", stacks)
	stacks.Dequeue()
	fmt.Printf("Stacks => %+v\n", stacks)
	stacks.Dequeue()
	fmt.Printf("Stacks => %+v\n", stacks)
}

func (s *stack) Push(item int) {
	*s = append(*s, item)
}

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Pop() {
	if len(*s) == 0 {
		s.IsEmpty()
	} else {
		index := len(*s) - 1
		pop := (*s)[index]
		fmt.Printf("%d popped \n", pop)
		*s = (*s)[:index]
	}
}

func (s *stack) Dequeue() {
	if len(*s) == 0 {
		s.IsEmpty()
	} else {
		item := (*s)[0]
		fmt.Println("Dequeued Element", item)
		*s = (*s)[1:len(*s)]
	}
}

func (s *stack) Enqueue(item int) {
	*s = append(*s, item)
}
