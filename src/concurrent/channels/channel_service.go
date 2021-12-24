package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
)

// Channels as a handle on a service

func main() {
	joe := service.Boring("Joe")
	ann := service.Boring("Ann")
	for i := 0; i < 5; i++ {
		fmt.Println(<-joe)
		fmt.Println(<-ann)
	}
	fmt.Println("You're both boring; I'm leaving")
}
