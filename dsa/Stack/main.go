package main

import "fmt"

type Stack struct {
	items []int
}

// Push
func (s *Stack) Push(i int) {

	s.items = append(s.items, i)
}

// Pop
func (s *Stack) Pop() int {

	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {

	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(30)
	myStack.Push(550)
	myStack.Push(5540)
	myStack.Push(3043)
	myStack.Push(303)
	myStack.Pop()
	fmt.Println(myStack)

}
