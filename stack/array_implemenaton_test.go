package stack

import "testing"

func TestArrayImplementation(t *testing.T) {
	stack := newArrayStack(5)
	err := push( stack,10)
	if err != nil {
		panic(err)
	}
}