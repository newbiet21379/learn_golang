package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
	"time"
)

func main() {
	c := service.Boring("Joe")
	timeout := time.After(5 * time.Second) // Time out after total 5 sec
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout: // After total time end loop
			fmt.Println("You are too slow! Time's up!!!")
			return
		case <-time.After(1 * time.Second): // No mess in 1 sec end
			fmt.Println("Be fast !!!")
			return
		}
	}
}
