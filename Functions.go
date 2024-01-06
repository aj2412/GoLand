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

func main() {
	price, no := 90, 10
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)

	area, perimeter := rectProps(10.8, 5.6)
	fmt.Printf("Area %f Perimeter %f", area, perimeter)

	a, p := rectPropsWithNamed(100.0, 50.0)
	fmt.Printf("\nArea %f Perimeter %f", a, p)
}
