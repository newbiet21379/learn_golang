package main

import (
	"Learn_Golang/src/interface"
	"Learn_Golang/src/struct"
	"fmt"
)
func main() {
	var taskInterface _interface.TaskInterface
	// interface {} can receive any Type
	taskInterface = _struct.Task{Capacity: 5, Total: 10} // Type Task implements Task Interface
	getTask, _ := taskInterface.(_struct.Task)
	fmt.Println("Shortest Combination:",taskInterface.FindShortestCombination(getTask))
	fmt.Println("Short SubString:",taskInterface.ShortestSubString("abaabaaabaa"))
	//fmt.Println("Absolute", _struct.MyFloat(-math.Sqrt2).Abs())
}

