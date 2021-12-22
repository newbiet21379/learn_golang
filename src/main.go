package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	_ "sync"
	"time"
)

type Rating struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}
type Product struct {
	Id          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float32 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rate        Rating  `json:"rating"`
}

func main() {
	started := time.Now()
	urlAll := "https://fakestoreapi.com/products"
	//ctx , cancel := context.WithCancel(context.Background())
	//var wg sync.WaitGroup

	var (
		products []Product
		urls     []string
	)

	json.Unmarshal([]byte(fetch(urlAll)), &products)
	productChannel := make(chan Product)
	done := make(chan int)
	//wg.Add(len(products))

	for _, item := range products {
		urls = append(urls, urlAll+"/"+strconv.Itoa(item.Id))
	}

	for _, url := range urls {
		go mapFromURL(url, done, productChannel)
	}

	//wg.Wait()
	//close(productChannel)
	for i := len(urls); i > 0; {
		select {
		case v := <-productChannel:
			fmt.Printf("From Channel :%+v\n", v)
		case <-done:
			i--
		}
	}

	fmt.Printf("Done in %v \n", time.Since(started))

}

func mapFromURL(url string, done chan<- int, q chan<- Product) {
	var product Product
	res := fetch(url)
	err := json.Unmarshal([]byte(res), &product)
	if err != nil {
		return
	}
	q <- product
	done <- 1
}

func fetch(url string) string {
	productClient := http.Client{
		Timeout: 2 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "Something is wrong"
	}
	res, getErr := productClient.Do(req)
	if getErr != nil {
		return "Something is wrong"
	}
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "Something is wrong"
	}

	return string(body)
}
