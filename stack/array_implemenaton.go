package stack

import (
	"errors"
)

// Generic type structure to create stack using simple array
type ArrayStack struct {
	top      int
	capacity int
	arr      []interface{}
}

// This will create a new stack
func newArrayStack(capacity int) *ArrayStack {
	var stack ArrayStack
	stack.top = -1
	stack.capacity = capacity
	return &stack
}

// This will check if the stack is empty
func isEmtpty(stack *ArrayStack) bool {
	return stack.top == -1
}

// This will check if the stack is full
func isFull(stack *ArrayStack) bool {
	return stack.top == stack.capacity-1
}

// this will return to top element in the stack
func top(stack *ArrayStack) interface{}{
	return stack.arr[stack.top]
}

// this will push an element into stack
func push(stack *ArrayStack, a interface{}) error {
	if isFull(stack) {
		return errors.New("stack is full")
	} else {
		stack.arr = append(stack.arr, a)
		stack.top++
	}
	return nil
}

func pop(stack *ArrayStack) (interface{}, error) {
	var element interface{}
	if isEmtpty(stack) {
		return nil, errors.New("stack is empty")
	} else {
		element = stack.arr[stack.top]
		stack.top--
	}
	return element, nil
}
