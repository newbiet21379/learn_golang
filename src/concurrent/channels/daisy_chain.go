package main

import (
	"fmt"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right // Left channel receive right value with +1
}

// Golang's goroutines is extremely lightweight compare to normal thread
func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	now := time.Now()
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}
	go func(c chan int) {
		c <- 1
	}(right)
	fmt.Println(<-leftmost)
	fmt.Printf("Time from start : %v", time.Since(now).Seconds())
}
