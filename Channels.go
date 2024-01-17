package main

import (
	"fmt"
)

func hello(done chan bool) {
	fmt.Println("hello go routine is going to sleep")
	//time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func producer(chnl chan int) {
	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go goroutine")
	go hello(done)
	a := <-done
	//time.Sleep(2 * time.Second)
	fmt.Println("Main received data is", a)

	// *********** ------------ ************

	number := 589
	sqrch := make(chan int)
	go calcSquares(number, sqrch)
	squares := <-sqrch
	fmt.Println("Final output", squares)

	// *********** ------------ ************

	ch := make(chan int)
	go producer(ch)
	for v := range ch {
		//time.Sleep(1 * time.Second)
		fmt.Println("Received ", v)
	}

}
