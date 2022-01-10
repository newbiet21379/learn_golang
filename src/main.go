package main

//
//import (
//	"encoding/json"
//	"fmt"
//	"io"
//	"io/ioutil"
//	"net/http"
//	"strconv"
//	_ "sync"
//	"time"
//)
//
//type Rating struct {
//	Rate  float32 `json:"rate"`
//	Count int     `json:"count"`
//}
//type Product struct {
//	Id          int     `json:"id"`
//	Title       string  `json:"title"`
//	Price       float32 `json:"price"`
//	Description string  `json:"description"`
//	Category    string  `json:"category"`
//	Image       string  `json:"image"`
//	Rate        Rating  `json:"rating"`
//}
//
//func main() {
//	started := time.Now()
//	urlAll := "https://fakestoreapi.com/products"
//
//	var (
//		products []Product
//		urls     []string
//		//fetchList []Product
//	)
//
//	err := json.Unmarshal([]byte(fetch(urlAll)), &products)
//	if err != nil {
//		return
//	}
//	productChannel := make(chan Product)
//	done := make(chan int)
//
//	for _, item := range products {
//		urls = append(urls, urlAll+"/"+strconv.Itoa(item.Id))
//	}
//
//	for _, url := range urls {
//		//mapFromURLSync(url, &fetchList)
//		go mapFromURL(url, done, productChannel)
//	}
//
//	//for i, p := range fetchList {
//	//	fmt.Printf("Product index : %d and Value :%+v", i, p)
//	//}
//	for i := len(urls); i > 0; {
//		select {
//		case v := <-productChannel:
//			fmt.Printf("From Channel :%+v\n", v)
//		case <-done:
//			i--
//		}
//	}
//
//	fmt.Printf("Done in %v \n", time.Since(started))
//
//}
//
//func mapFromURLSync(url string, products *[]Product) {
//	var product Product
//	res := fetch(url)
//	err := json.Unmarshal([]byte(res), &product)
//	if err != nil {
//		return
//	}
//	*products = append(*products, product)
//}
//
//func mapFromURL(url string, done chan<- int, q chan<- Product) {
//	var product Product
//	res := fetch(url)
//	err := json.Unmarshal([]byte(res), &product)
//	if err != nil {
//		return
//	}
//	q <- product
//	done <- 1
//}
//
//func fetch(url string) string {
//	productClient := http.Client{
//		Timeout: 2 * time.Second,
//	}
//
//	req, err := http.NewRequest(http.MethodGet, url, nil)
//	if err != nil {
//		return "Something is wrong"
//	}
//	res, getErr := productClient.Do(req)
//	if getErr != nil {
//		return "Something is wrong"
//	}
//	if res.Body != nil {
//		defer func(Body io.ReadCloser) {
//			err := Body.Close()
//			if err != nil {
//
//			}
//		}(res.Body)
//	}
//
//	body, readErr := ioutil.ReadAll(res.Body)
//	if readErr != nil {
//		return "Something is wrong"
//	}
//
//	return string(body)
//}
