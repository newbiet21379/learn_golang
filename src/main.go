package main
import (
	"Learn_Golang/src/struct"
	"Learn_Golang/src/learn"
	"fmt"
	"math"
)
func main() {
	t := learn.Task{}
	fmt.Println("Shortest Combination:",t.FindShortestCombination(5,10))
	fmt.Println("Short SubString:",t.ShortestSubString("abaabaaabaa"))
	fmt.Println("Absolute", _struct.MyFloat(-math.Sqrt2).Abs())
	var abser _struct.Abser
	float := _struct.MyFloat(-math.Sqrt2)
	vertex := _struct.Vertex{3, 4}

	abser = float  // a MyFloat implements Abser
	abser = &vertex // a *Vertex implements Abser

	fmt.Println(abser.Abs())
}

