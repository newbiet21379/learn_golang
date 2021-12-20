package functions

import (
	"fmt"
	"strconv"
	"time"
)
//func add(x, y int) int {
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	var (
		n int
		s string
	)
	fmt.Println(add(42, 13))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(split(17))
	for {
		if n < 10{
			break
		}
		if smt := n; smt > 5{
			fmt.Println("Something: ",smt)
		}
		fmt.Println("n: ",n)
		n++
	}

	for ;n<100;{
		s+= strconv.Itoa(n)
		n+= 5
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}

	defer println("After run")

}

