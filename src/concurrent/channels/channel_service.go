package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
	"time"
)

// Channels as a handle on a service

func main() {
	joe := service.Boring("Joe")
	ann := service.Boring("Ann")
	timeout := time.After(time.Duration(5 * time.Second))
	for {
		select {
		case s := <-joe:
			fmt.Println(s)
		case an := <-ann:
			fmt.Println(an)
		case <-timeout:
			fmt.Println("Times up ")
			return
		}

	}
	fmt.Println("You're both boring; I'm leaving")
}
