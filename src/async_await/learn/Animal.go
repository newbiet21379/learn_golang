package main

import (
	"fmt"
)

type Animal struct {
	Age int
	Leg int
}

func (a *Animal) Move() {
	fmt.Println("An animal moves by ", a.Leg, " legs.")
}

func main() {

	fmt.Print("=====================================================================\n")
	sliceWithCap := make([]int, 0, 5)
	fmt.Printf("Length of slices with cap: %d , %d\n", len(sliceWithCap), cap(sliceWithCap))
	fmt.Print("=====================================================================\n")
	sliceWithoutCap := make([]int, 5)
	fmt.Printf("Length of slices without cap: %d , %d", len(sliceWithoutCap), cap(sliceWithoutCap))

	word := "New To This"

	for index, char := range word {
		fmt.Printf("Index: %d\t Value: %d\n", index, char)
	}

	for _, char := range word {
		fmt.Printf("Value: %d\n", char)
	}

	for i := 0; i < len(word); i++ {
		fmt.Printf("Index: %d\t Value: %d\n", i, word[i])
	}
}
