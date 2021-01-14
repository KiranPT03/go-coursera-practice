package main

import "fmt"

func product(a int, b int, c chan int) {
	product := a * b
	c <- product
}

func main() {
	c := make(chan int)
	go product(3, 4, c)
	go product(5, 6, c)

	x := <-c
	y := <-c

	totalProduct := x * y
	fmt.Println(totalProduct)

}
