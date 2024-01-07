package main

import "fmt"

type employee struct {
	salary  int
	country string
}

func main() {
	uno := make(map[string]float64)
	uno["Arjit"] = 8.8
	uno["Ashu"] = 9.1
	uno["Naman"] = 9.3
	fmt.Println("Length of", uno, "is", len(uno))

	newn := "Ashu"
	value, ok := uno[newn]
	if ok == true {
		fmt.Println("Marks of", newn, "are", value)
	} else {
		fmt.Println("Not found.")
	}

	for key, value := range uno {
		fmt.Printf("Marks of %s is %v\n", key, value)
	}

	delete(uno, "Arjit")
	fmt.Println(uno)

	// MAPS OF STRUCT

	e1 := employee{
		salary:  12000,
		country: "USA",
	}
	e2 := employee{
		salary:  14000,
		country: "Canada",
	}
	e3 := employee{
		salary:  15000,
		country: "India",
	}
	employeeInfo := map[string]employee{
		"Arjit":  e1,
		"Ashu":   e2,
		"Naman ": e3,
	}

	for name, info := range employeeInfo {
		fmt.Printf("Employee: %s Salary: $%d  Country: %s\n", name, info.salary, info.country)
	}

}
