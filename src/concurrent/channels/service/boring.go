package service

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/entity"
	"math/rand"
	"time"
)

func Boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func BoringMessage(msg entity.Message) <-chan entity.Message {
	c := make(chan entity.Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- entity.Message{Str: fmt.Sprintf("%s : %d", msg.Str, i), Wait: waitForIt}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
			<-waitForIt
		}
	}()
	return c
}

// Using c as the channel to listen to whoever ready to talk between the 2 input
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			c <- <-input1 // Getting input from channel Input1
		}
	}()
	go func() {
		for {
			c <- <-input2 // Getting input from channel Input2
		}
	}()
	return c
}

// Using c as the channel to listen to whoever ready to talk between the 2 input
func FanInMessaging(input1, input2 <-chan entity.Message) <-chan entity.Message {
	c := make(chan entity.Message)
	go func() {
		for {
			c <- <-input1 // Getting input from channel Input1
		}
	}()
	go func() {
		for {
			c <- <-input2 // Getting input from channel Input2
		}
	}()
	return c
}
