package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/interface"
	"github.com/newbiet21379/learn_golang/src/struct"
	"math"
)
func main() {
	var capacity int
	fmt.Print("Input total numbers of cups: ")
	_, err := fmt.Scanf("%d",&capacity)
	if err != nil {
		return
	}
	_, err = fmt.Scanln()
	if err != nil {
		return 
	}
	var total int
	fmt.Print("Input total amount of water: ")
	_, err = fmt.Scanf("%d",&total)
	if err != nil {
		return
	}
	fmt.Println(capacity,total)
	var taskInterface _interface.TaskInterface
	// interface {} can receive any Type
	taskInterface = _struct.Task{Capacity: capacity, Total: total} // Type Task implements Task Interface
	getTask, _ := taskInterface.(_struct.Task)
	fmt.Println("Shortest Combination:",taskInterface.FindShortestCombination(getTask))
	fmt.Println("Short SubString:",taskInterface.ShortestSubString("abaabaaabaa"))
	fmt.Println("Absolute", _interface.MyFloat(-math.Sqrt2).Abs())
	_, err = fmt.Scanln()
	if err != nil {
		return
	}
}

