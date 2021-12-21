package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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
	url := "https://fakestoreapi.com/products"

	//var wg sync.WaitGroup
	//wg.Add(1)

	var products []Product
	var temp Product
	json.Unmarshal([]byte(fetch(url)), &products)
	fetchProduct := make(chan Product, len(products))
	for _, item := range products {
		go func(item Product) {

			err := json.Unmarshal([]byte(fetch(url+"/"+strconv.Itoa(item.Id))), &temp)
			if err != nil {
				return
			}
			fetchProduct <- temp
		}(item)
	}
	close(fetchProduct)
	//wg.Wait()
	for i := range fetchProduct {
		fmt.Printf("Product :%#v\n", i)
	}
	fmt.Printf("Done in %v \n", time.Since(started))

}

func fetch(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		return "Something is wrong"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Something is wrong"
	}
	resp.Body.Close()
	return string(body)
}
