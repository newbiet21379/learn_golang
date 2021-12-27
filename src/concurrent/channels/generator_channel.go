package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
)

func main() {
	c := service.Boring("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c) // Receive expression is just a value
	}
	fmt.Println("You're boring; I'm leaving")
}
