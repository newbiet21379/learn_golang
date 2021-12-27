package Promise_race

import (
	"fmt"
	"math/rand"
	"time"
)

// JavaScript
// ---

//const one = async () => {
// simulate a workload
//sleep(Math.floor(Math.random() * Math.floor(2000)));
//return 1;
//};
//
//const two = async () => {
// simulate a workload
//sleep(Math.floor(Math.random() * Math.floor(1000)));
//sleep(Math.floor(Math.random() * Math.floor(1000)));
//return 2;
//};
//
//const r = await Promise.race(one(), two());
//console.log(r);

// Go
// ---

func one() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// simulate a workload
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(2000)))
		r <- 1
	}()

	return r
}

func two() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// simulate a workload
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		time.Sleep(time.Millisecond * time.Duration(rand.Int63n(1000)))
		r <- 2
	}()

	return r
}

func main() {
	var r int32
	select {
	case r = <-one():
	case r = <-two():
	}

	fmt.Println(r)
}