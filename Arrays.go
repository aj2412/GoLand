package main

import "fmt"

func changeLocal(num [3]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
	return
}

func main() {
	var arr1 [3]int //int array with length 3
	arr1[0] = 12
	arr1[1] = 78
	arr1[2] = 50
	fmt.Println(arr1)

	arr2 := [3]int{1, 7, 5} // shorthand declaration to create array
	fmt.Println(arr2)

	arr3 := [...]string{"Germany", "Venice", "Spain"}
	arr4 := arr3
	arr4[0] = "India"
	fmt.Println("Length of", arr4, "is", len(arr4))

	changeLocal(arr2)
	fmt.Println(arr2, "- After function execution")

	var sum int = 0
	for _, v := range arr2 {
		sum = sum + v
	}
	fmt.Println("Sum of", arr2, "is", sum)

	arr5 := [2][2]string{
		{"add", "sub"},
		{"mul", "div"},
	}
	fmt.Println(arr5)
	for _, v := range arr5 {
		for _, s := range v {
			fmt.Printf("%s ", s)
		}
		fmt.Println()
	}

	c := []int{6, 7, 8} //creates an array and returns a slice reference
	fmt.Println(c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Println("array before", darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
	fmt.Printf("length of slice %d capacity %d", len(dslice), cap(dslice))

	veg := []string{"cauli", "ladyfinger"}
	fruits := []string{"apple", "mango", "grapes"}
	food := append(veg, fruits...)
	fmt.Println("Eatable items are", food)

	pls := [3][2]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for i, _ := range v1 {
			fmt.Printf("%s ", v1[i])
		}
		fmt.Printf("\n")
	}

	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	countriesCpy := make([]string, len(neededCountries))
	fmt.Println(countriesCpy)
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Println(countriesCpy)

}
