package stack

import (
	"fmt"

	"github.com/abstract300/ds/linkedlists"
)

type StackDLL struct {
	top  int
	size int
	list *linkedlists.DoubleList
}

func NewStackDLL(size int) *StackDLL {
	var stack StackDLL
	stack.list = linkedlists.MakeList()
	stack.size = size
	return &stack
}

func (sd *StackDLL) Push(data interface{}) error {
	if sd.isFull() {
		return fmt.Errorf("%s", "Stack if full")
	}

	sd.list.AddFront(data)
	sd.top++
	return nil
}

func (sd *StackDLL) Pop() (interface{}, error) {

	if sd.isEmpty() {
		return nil, fmt.Errorf("Stack is empty")
	}

	data := sd.list.PeekFront()
	sd.list.Remove(0)
	sd.top--

	return data, nil
}

func (sd *StackDLL) Peek() (interface{}, error) {
	if sd.isEmpty() {
		var data interface{}
		return data, fmt.Errorf("%s", "Stack is empty")
	}
	return sd.list.PeekFront(), nil
}

func (sd *StackDLL) isFull() bool {
	if sd.top == sd.size {
		return true
	}
	return false
}

func (sd *StackDLL) isEmpty() bool {
	if sd.top == 0 {
		return true
	}
	return false
}
