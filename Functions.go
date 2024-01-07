package main

import (
	"fmt"
)

func calculateBill(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

func rectPropsWithNamed(length, width float64) (a, p float64) {
	a = length * width
	p = (length + width) * 2
	return //no explicit return value
}

func find(num int, nums ...int) { //variadic function
	fmt.Printf("Type of nums is %T\n", nums)
	found := false
	for i, v := range nums {
		if num == v {
			found = true
			fmt.Printf("The number %v is found at index number %v\n", num, i)
		}
	}
	if !found {
		fmt.Printf("Number is not found.\n")
	}
	return
}

func change(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s, "- inside function")
}

func main() {
	price, no := 90, 10
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f\n", area, perimeter)

	a, p := rectPropsWithNamed(100.0, 50.0)
	fmt.Printf("Area %f Perimeter %f\n", a, p)

	find(12, 3, 4, 5, 6, 7, 10, 12, 12, 21)

	nums := []int{89, 90, 95}
	find(95, nums...) // way of passing slice to variadic function

	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome, "- after function")

}
