package single_async

import (
	"fmt"
	"math/rand"
	"time"
)

//JavaScript
//
//---
//
//const longRunningTask = async () => {
//
////simulate a workload
//
//sleep(3000);
//
//return Math.floor(Math.random() * Math.floor(100));
//};
//
//const r = await longRunningTask();
//console.log(r);

func LongRunningTask() <-chan int32 {
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
	r := <-LongRunningTask()
	fmt.Println(r)
}
