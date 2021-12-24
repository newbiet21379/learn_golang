package main

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
	"math/rand"
	"time"
)

//No locks. No condition variables. No callbacks
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := service.GoogleVer3(
		"golang",
	)
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
