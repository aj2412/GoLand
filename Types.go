package main

import (
	"fmt"
	"unsafe"
)

func main() {

	a := true
	b := false
	fmt.Println("a:", a, "b:", b)
	c := a && b
	fmt.Println("c:", c)
	d := a || b
	fmt.Println("d:", d)

	var x int = 89
	y := 95
	fmt.Println("Value of x is", x, "and y is", y)
	fmt.Printf("Type of x is %T, Size of x is %d", x, unsafe.Sizeof(x))
	fmt.Printf("\nType of y is %T, Size of y is %d", y, unsafe.Sizeof(y))

	no1, no2 := 56, 89
	fmt.Println("sum", no1+no2, "diff", no1-no2)

	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	i := 10
	j := 11.3
	sum := i + int(j)
	fmt.Println("sum:", sum)

	var k float64 = float64(i)
	fmt.Println("k:", k)
}
