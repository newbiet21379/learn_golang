package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/newbiet21379/learn_golang/src/concurrent/channels/entity"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
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

//Function read file with input and output channels
//Input: file name
//Output: channel of string
func readFile(fileName string) chan string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c
}

// Function to write to file from channel
// Input: file name, channel of string
// Output: nil
func writeFile(fileName string, c chan string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for line := range c {
		fmt.Fprintln(file, line)
	}
}

// FanIn Function with input channels
// Input: channels of string
// Output: channel of string
func fanIn(cs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range cs {
		go func(ch chan string) {
			for msg := range ch {
				c <- msg
			}
		}(ch)
	}
	return c
}

//Fanout function with input channel
//Input: channel of string
//Output: channels of string
func fanOut(c chan string) (chan string, chan string) {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for msg := range c {
			c1 <- msg
			c2 <- msg
		}
	}()
	return c1, c2
}
