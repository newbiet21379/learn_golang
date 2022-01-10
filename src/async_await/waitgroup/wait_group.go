package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
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
	var wg sync.WaitGroup

	var (
		products  []Product
		urls      []string
		fetchList []Product
	)

	err := json.Unmarshal([]byte(fetch(urlAll)), &products)
	if err != nil {
		return
	}
	wg.Add(len(products))

	for _, item := range products {
		urls = append(urls, urlAll+"/"+strconv.Itoa(item.Id))
	}

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			mapFromURL(url, &fetchList)
		}(url)
	}

	wg.Wait()
	//close(productChannel)
	for i, p := range fetchList {
		fmt.Printf("From Wait Group %d :%+v\n", i, p)
	}

	fmt.Printf("Done in %v \n", time.Since(started))

}

func mapFromURL(url string, products *[]Product) {
	var product Product
	res := fetch(url)
	err := json.Unmarshal([]byte(res), &product)
	if err != nil {
		return
	}
	*products = append(*products, product)
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
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {

			}
		}(res.Body)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return "Something is wrong"
	}

	return string(body)
}
