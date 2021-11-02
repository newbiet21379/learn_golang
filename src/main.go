package main

import (
	"Learn_Golang/src/learn"
	"Learn_Golang/src/struct"
	"fmt"
	"math"
)
func main() {
	var taskInterface _struct.TaskInterface

	taskInterface = learn.Task{Capacity: 5, Total: 10} // Type Task implements Task Interface
	getTask := taskInterface.(learn.Task)
	fmt.Println("Shortest Combination:",taskInterface.FindShortestCombination(getTask))
	fmt.Println("Short SubString:",taskInterface.ShortestSubString("abaabaaabaa"))
	fmt.Println("Absolute", _struct.MyFloat(-math.Sqrt2).Abs())
}

