package _interface

import "testing"

func TestInterface(t *testing.T) {
	var b Vehicle
	b = &Car{}
	b.Wheels()
}