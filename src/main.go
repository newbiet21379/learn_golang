package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

func main() {
	started := time.Now()
	foods := []string{"mushroom pizza", "pasta", "kebab", "cake"}
	foodChan := make(chan string)
	//var wg sync.WaitGroup
	//wg.Add(len(foods))
	for _, food := range foods {
		//cook(food)
		go func(food string) {
			foodChan <- cook(food)
			//wg.Done()
		}(food)
	}
	//wg.Wait()
	for i := 0; i < len(foods); i++ {
		fmt.Printf("Channel value : %v\n", <-foodChan)
	}
	close(foodChan)
	fmt.Printf("Done in %v\n", time.Since(started))
}

func cook(food string) string {
	fmt.Printf("cooking %s...\n", food)
	time.Sleep(2 * time.Second)
	return fmt.Sprintf("done cooking %s\n", food)
}
