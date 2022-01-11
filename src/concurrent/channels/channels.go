package main

import "fmt"

// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// Sender and receiver must both be ready to play their part
// in the communication. Otherwise, we wait until they do
func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) //
	go sum(s[len(s)/2:], c)
	// Wait for value to be sent
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// waits for receiver to be ready
	c <- sum // send sum to c
}
