package service

import (
	"fmt"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/entity"
	"math/rand"
	"time"
)

var (
	Web    = FakeSearch("web")
	Image  = FakeSearch("image")
	Video  = FakeSearch("video")
	Web1   = FakeSearch("web 1")
	Image1 = FakeSearch("image 1")
	Video1 = FakeSearch("video 1")
	Web2   = FakeSearch("web 2")
	Image2 = FakeSearch("image 2")
	Video2 = FakeSearch("video 2")
)

// Google Traditional synchronization
func Google(query string) (results []entity.Result) {
	results = append(results, Web(query))
	results = append(results, Image(query))
	results = append(results, Video(query))
	return
}

// GoogleConcurrent Wait for all results
func GoogleConcurrent(query string) (results []entity.Result) {
	c := make(chan entity.Result)
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()
	for i := 0; i < 3; i++ {
		result := <-c
		results = append(results, result)
	}
	return
}

func GoogleVer2(query string) (results []entity.Result) {
	c := make(chan entity.Result)
	go func() {
		c <- Web(query)
	}()
	go func() {
		c <- Image(query)
	}()
	go func() {
		c <- Video(query)
	}()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Time out !!!")
			return
		}
	}
	return
}

func GoogleVer3(query string) (results []entity.Result) {
	c := make(chan entity.Result)
	go func() {
		c <- First(query, Web1, Web2)
	}()
	go func() {
		c <- First(query, Image1, Image2)
	}()
	go func() {
		c <- First(query, Video1, Video2)
	}()
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 3; i++ {
		select {
		case result := <-c:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Time out !!!")
			return
		}
	}
	return
}

// First Avoid Timeout ( Discarding results from slow server
func First(query string, replicas ...entity.Search) entity.Result {
	c := make(chan entity.Result)
	searchReplica := func(i int) {
		c <- replicas[i](query)
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-c
}

func FakeSearch(kind string) entity.Search {
	return func(query string) entity.Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return entity.Result{Data: fmt.Sprintf("%s result for %q\n", kind, query)}
	}
}
