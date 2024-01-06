package main

import (
	"fmt"
	"math"
)

func main() {

	var myage int
	fmt.Println("My age is ", myage)
	myage = 21
	fmt.Println("My age is ", myage)
	myage = 100
	fmt.Println("My age is ", myage)

	var (
		name   = "Arjit"
		age    = 21
		height int
	)

	fmt.Println("My name is", name, ", Age is", age, "and Height is", height)

	count := 10
	fmt.Println("Count is ", count)

	a, b := 10.1, 20.0
	c := math.Min(a, b)
	fmt.Println("Minimum of a and b is ", c)
}
