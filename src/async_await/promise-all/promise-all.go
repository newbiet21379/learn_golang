package promise_all

// JavaScript
// ---

//const longRunningTask = async () => {
// simulate a workload
//sleep(3000);
//return Math.floor(Math.random() * Math.floor(100));
//};
//
//const [a, b, c] = await Promise.all(
//longRunningTask(),
//longRunningTask(),
//longRunningTask()
//);
//console.log(a, b, c);

import (
	"fmt"
	"math/rand"
	"time"
)

func longRunningTask() <-chan int32 {
	r := make(chan int32)

	go func() {
		defer close(r)

		// simulate a workload
		time.Sleep(time.Second * 3)
		r <- rand.Int31n(100)
	}()

	return r
}

func main() {
	aCh, bCh, cCh := longRunningTask(), longRunningTask(), longRunningTask()
	a, b, c := <-aCh, <-bCh, <-cCh

	fmt.Println(a, b, c)
}
