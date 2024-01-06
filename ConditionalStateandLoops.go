package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num1 := 10
	if num1%3 == 0 {
		fmt.Println("The remainder is zero.")
	} else if num1%3 == 1 {
		fmt.Println("The remainder is one.")
	} else {
		fmt.Println("The remainder is two.")
	}

	num2 := 99
	if num2 <= 50 {
		fmt.Println(num2, "is less than or equal to 50")
	} else if num2 >= 51 && num2 <= 100 {
		fmt.Println(num2, "is between 51 and 100")
	} else {
		fmt.Println(num2, "is greater than 100")
	}

	//for statement

	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}

	fmt.Println()

	for i := 1; i <= 10; i++ {
		if i > 6 {
			break //loop is terminated if i > 5
		}
		fmt.Print(i, " ")
	}
	fmt.Println("- line after for loop")

	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Print(i, " ")
	}
	fmt.Println("- line of odd numbers")

	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}

	fmt.Println()

	//switch statement

	switch num := 4; {
	case num < 50:
		if num < 0 {
			break
		}
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Println(num, "is lesser than 200\n")
	}

	//this the format to use label of for loop

randomloop:
	for {
		switch i := rand.Intn(100); {
		case i%2 == 0:
			fmt.Println("Generated even number", i)
			break randomloop
		default:
			fmt.Println("Wrong number generated", i)
		}
	}
}
