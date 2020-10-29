package _interface

import "fmt"

type Vehicle interface {
	Wheels()
}

type Car struct{}

type Bike struct{}

func (c *Car) Wheels() {
	fmt.Print("4 Wheels")
}

func (b *Bike) Wheels() {
	fmt.Printf("2 Wheels")
}
