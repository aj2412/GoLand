package main

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

type add struct {
	city  string
	state string
}

type Person struct {
	string
	int
	address add
}

func main() {
	// *********** ------------ ************
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}

	// *********** ------------ ************

	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	// *********** ------------ ************

	emp3 := struct {
		firstName string
		lastName  string
		age       int
		salary    int
	}{
		firstName: "Andreah",
		lastName:  "Nikola",
		age:       31,
		salary:    5000,
	}
	fmt.Println("Employee 3", emp3)

	// *********** ------------ ************

	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	emp8a := &emp8
	fmt.Println("First Name:", (*emp8a).firstName)
	fmt.Println("Age:", (*emp8a).age)

	// *********** ------------ ************

	p1 := Person{
		string: "Arjit",
		int:    21,
	}
	fmt.Println("Name of the person is", p1.string)
	fmt.Println("Age of the person is", p1.int)

	// *********** ------------ ************

	p := Person{
		string: "Ashu",
		int:    20,
		address: add{
			city:  "Hisar",
			state: "Haryana",
		},
	}
	fmt.Println("Name:", p.string)
	fmt.Println("Age:", p.int)
	fmt.Println("City:", p.address.city)
	fmt.Println("State:", p.address.state)

	// *********** ------------ ************

	p2 := Person{
		string: "A",
		int:    21,
	}
	p3 := Person{
		string: "A",
		int:    24,
	}
	if p2 == p3 {
		fmt.Println("They are equal.")
	} else {
		fmt.Println("Not equal.")
	}

}
