package main

import "fmt"

func main() {
	c := make(chan int, 5)
	c <- 1
	c <- 5
	c <- 4
	c <- 8
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}