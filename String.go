package main

import (
	"fmt"
	"unicode/utf8"
)

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
}

func printChars(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, char := range s {
		fmt.Printf("%c lies at byte %d\n", char, index)
	}
}

//func mutate(s string)string {
//	s[0] = 'a'//any valid unicode character within single quote is a rune
//	return s
//}

func main() {
	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n\n")

	name = "Arjit"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n")
	charsAndBytePosition(name)

	fmt.Println()

	byteSlice := []byte{67, 97, 102, 195, 169} //ASCII values stored in slices.
	str := string(byteSlice)
	fmt.Println(str)

	fmt.Printf("Length: %d\n", utf8.RuneCountInString(str))
	fmt.Printf("Number of bytes: %d\n", len(str))

	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)
	final := string1 + " " + string2
	fmt.Println(final)

	//mutate(final)  this function will give an error as a string in GO is immutable.

}
