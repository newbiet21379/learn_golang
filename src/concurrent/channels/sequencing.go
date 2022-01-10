package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/entity"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
	"time"
)

func main() {
	ann := new(entity.Message)
	ann.Str = "Ann"
	joe := new(entity.Message)
	joe.Str = "Joe"
	start := time.Now()
	c := service.FanInMessaging(service.BoringMessage(*ann), service.BoringMessage(*joe))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.Str)
		msg2 := <-c
		fmt.Println(msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
	fmt.Println("You're both boring; I'm leaving")
	fmt.Printf("Total time is :%v", time.Since(start))
}
