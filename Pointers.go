package main

import (
	"fmt"
)

func modify(arr *[3]int) {
	(*arr)[0] = 101
}

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	v1 := 225
	var v2 *int
	fmt.Println("\n", v2)
	v2 = &v1
	fmt.Println("After Initializing the value becomes ", v2)

	fmt.Printf("\nValue is %d, Type is %T, Address is %v\n", *v2, v2, v2)

	arr := [3]int{89, 90, 91}
	modify(&arr)
	fmt.Println(arr)

}
