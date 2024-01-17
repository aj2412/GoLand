package main

import "fmt"

type Desc interface {
	Describe()
}
type Human struct {
	name string
	age  int
}

func (p Human) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

type Address struct {
	state   string
	country string
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("State %s Country %s", a.state, a.country)
}

type SalCalculator interface {
	DisplaySalary()
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type EmployeeOperations interface {
	SalCalculator
	LeaveCalculator
}

type Worker struct {
	firstName   string
	lastName    string
	basicPay    int
	pf          int
	totalLeaves int
	leavesTaken int
}

func (e Worker) DisplaySalary() {
	fmt.Printf("%s %s has salary $%d", e.firstName, e.lastName, (e.basicPay + e.pf))
}

func (e Worker) CalculateLeavesLeft() int {
	return e.totalLeaves - e.leavesTaken
}

func main() {
	var d1 Desc
	p1 := Human{"Sam", 25}
	d1 = p1
	d1.Describe()
	p2 := Human{"James", 32}
	d1 = &p2
	d1.Describe()

	var d2 Desc
	a := Address{"Washington", "USA"}
	// (Describe method has pointer receiver)compilation error if the following line is uncommented
	//d2 = a
	d2 = &a
	d2.Describe()

	// *********** ------------ ************

	e := Worker{
		firstName:   "Naveen",
		lastName:    "Ramanathan",
		basicPay:    5000,
		pf:          200,
		totalLeaves: 30,
		leavesTaken: 5,
	}
	var empOp EmployeeOperations = e
	empOp.DisplaySalary()
	fmt.Println("\nLeaves left =", empOp.CalculateLeavesLeft())

}
