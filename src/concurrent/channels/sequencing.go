package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/newbiet21379/learn_golang/src/concurrent/channels/entity"
	"github.com/newbiet21379/learn_golang/src/concurrent/channels/service"
)

func main() {
	ann := new(entity.Message)
	ann.Str = "Ann"
	joe := new(entity.Message)
	joe.Str = "Joe"
	start := time.Now()
	c := service.FanInMessaging(service.BoringMessage(*ann), service.BoringMessage(*joe))
	for i := 0; i < 5; i++ {
		msg1 := <-c
		fmt.Println(msg1.Str)
		msg2 := <-c
		fmt.Println(msg2.Str)
		msg1.Wait <- true
		msg2.Wait <- true
	}
	fmt.Println("You're both boring; I'm leaving")
	fmt.Printf("Total time is :%v", time.Since(start))
}

//Function read file with input and output channels
//Input: file name
//Output: channel of string
func readFile(fileName string) chan string {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	c := make(chan string)
	go func() {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			c <- scanner.Text()
		}
		close(c)
	}()
	return c
}

// Function to write to file from channel
// Input: file name, channel of string
// Output: nil
func writeFile(fileName string, c chan string) {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	for line := range c {
		fmt.Fprintln(file, line)
	}
}

// FanIn Function with input channels
// Input: channels of string
// Output: channel of string
func fanIn(cs ...chan string) chan string {
	c := make(chan string)
	for _, ch := range cs {
		go func(ch chan string) {
			for msg := range ch {
				c <- msg
			}
		}(ch)
	}
	return c
}

//Fanout function with input channel
//Input: channel of string
//Output: channels of string
func fanOut(c chan string) (chan string, chan string) {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		for msg := range c {
			c1 <- msg
			c2 <- msg
		}
	}()
	return c1, c2
}

// Find duplicate char in string
// Input: string
// Output: string
func findDuplicateChar(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		count := 0
		for j := 0; j < len(str); j++ {
			if str[i] == str[j] {
				count++
			}
		}
		if count > 1 {
			result += string(str[i])
		}
	}
	return result
}

// Replace char with * in string
// Input: string
// Output: string
func replaceChar(str string) string {
	var result string
	for i := 0; i < len(str); i++ {
		result += "*"
	}
	return result
}

// Find an element from slice of interface
// Input: slice of interface, element
// Output: bool
func findElement(slice []interface{}, element interface{}) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

// Find an element from slice of string
// Input: slice of string, element
// Output: bool
func findElementString(slice []string, element string) bool {
	for _, value := range slice {
		if value == element {
			return true
		}
	}
	return false
}

// Find an element from slice of interface and return struct match interface
// Input: slice of interface, element
// Output: struct
func findElementStruct(slice []interface{}, element interface{}) interface{} {
	for _, value := range slice {
		if value == element {
			return value
		}
	}
	return nil
}

// Parse Interface into struct
// Input: interface
// Output: struct
func parseInterface(i interface{}) interface{} {
	switch i.(type) {
	case int:
		return i.(int)
	case string:
		return i.(string)
	case float64:
		return i.(float64)
	case bool:
		return i.(bool)
	default:
		return nil
	}
}

type Book struct {
	Title  string
	Author string
	Price  float64
}

// Create array of interface from slice of struct
// Input: slice of struct
// Output: array of interface
func createArrayInterface(slice []Book) []interface{} {
	var result []interface{}
	for _, value := range slice {
		result = append(result, value)
	}
	return result
}

// Interface of Book Store
type BookStore interface {
	AddBook(book Book)
	GetBook(title string) Book
	GetAllBooks() []Book
	GetAllBooksByAuthor(author string) []Book
	GetAllBooksByPrice(price float64) []Book
	GetAllBooksByTitle(title string) []Book
	GetAllBooksByTitleAndAuthor(title string, author string) []Book
	GetAllBooksByTitleAndPrice(title string, price float64) []Book
	GetAllBooksByAuthorAndPrice(author string, price float64) []Book
	GetAllBooksByTitleAndAuthorAndPrice(title string, author string, price float64) []Book
}

// Implement Book Store interface
type BookStoreImpl struct {
	books []Book
}

// Book Store function AddBook
func (bs *BookStoreImpl) AddBook(book Book) {
	bs.books = append(bs.books, book)
}

// Book Store function GetBook
func (bs *BookStoreImpl) GetBook(title string) Book {
	for _, book := range bs.books {
		if book.Title == title {
			return book
		}
	}
	return Book{}
}

// Book Store function GetAllBooks
func (bs *BookStoreImpl) GetAllBooks() []Book {
	return bs.books
}

// Book Store function GetAllBooksByAuthor
func (bs *BookStoreImpl) GetAllBooksByAuthor(author string) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

// Book Store function GetAllBooksByPrice
func (bs *BookStoreImpl) GetAllBooksByPrice(price float64) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Price == price {
			result = append(result, book)
		}
	}
	return result
}

// Book Store function GetAllBooksByTitle
func (bs *BookStoreImpl) GetAllBooksByTitle(title string) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Title == title {
			result = append(result, book)
		}
	}
	return result
}

// Book Store function GetAllBooksByTitleAndAuthor
func (bs *BookStoreImpl) GetAllBooksByTitleAndAuthor(title string, author string) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Title == title && book.Author == author {
			result = append(result, book)
		}
	}
	return result
}

// Book Store get api GetAllBooksByTitleAndPrice
func (bs *BookStoreImpl) GetAllBooksByTitleAndPrice(title string, price float64) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Title == title && book.Price == price {
			result = append(result, book)
		}
	}
	return result
}

// Api for Book Store
func (bs *BookStoreImpl) GetAllBooksByAuthorAndPrice(author string, price float64) []Book {
	var result []Book
	for _, book := range bs.books {
		if book.Author == author && book.Price == price {
			result = append(result, book)
		}
	}
	return result
}

// create http server with BookStore
func createHttpServer(bookStore BookStore) {
	http.HandleFunc("/addBook", func(w http.ResponseWriter, r *http.Request) {
		var book Book
		json.NewDecoder(r.Body).Decode(&book)
		bookStore.AddBook(book)
	})
	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		book := bookStore.GetBook(title)
		json.NewEncoder(w).Encode(book)
	})
	http.HandleFunc("/getAllBooks", func(w http.ResponseWriter, r *http.Request) {
		books := bookStore.GetAllBooks()
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByAuthor", func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		books := bookStore.GetAllBooksByAuthor(author)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByPrice", func(w http.ResponseWriter, r *http.Request) {
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByPrice(price)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitle", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		books := bookStore.GetAllBooksByTitle(title)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitleAndAuthor", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")
		books := bookStore.GetAllBooksByTitleAndAuthor(title, author)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitleAndPrice", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByTitleAndPrice(title, price)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByAuthorAndPrice", func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByAuthorAndPrice(author, price)
		json.NewEncoder(w).Encode(books)
	})
	http.ListenAndServe(":8080", nil)
}

// implement BookStore interface for MongoDB
type MongoDB struct {
}

// AddBook function for MongoDB
func (mdb *MongoDB) AddBook(book Book) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test").C("books")
	err = collection.Insert(book)
	if err != nil {
		panic(err)
	}
}

// GetBook function for MongoDB
func (mdb *MongoDB) GetBook(title string) Book {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test").C("books")
	var result Book
	err = collection.Find(bson.M{"title": title}).One(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// GetAllBooks function for MongoDB
func (mdb *MongoDB) GetAllBooks() []Book {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test").C("books")
	var result []Book
	err = collection.Find(nil).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// GetAllBooksByAuthor function for MongoDB
func (mdb *MongoDB) GetAllBooksByAuthor(author string) []Book {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	collection := session.DB("test").C("books")
	var result []Book
	err = collection.Find(bson.M{"author": author}).All(&result)
	if err != nil {
		panic(err)
	}
	return result
}

// create MongoDB server with BookStore
func createMongoDBServer(bookStore BookStore) {
	http.HandleFunc("/addBook", func(w http.ResponseWriter, r *http.Request) {
		var book Book
		json.NewDecoder(r.Body).Decode(&book)
		bookStore.AddBook(book)
	})
	http.HandleFunc("/getBook", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		book := bookStore.GetBook(title)
		json.NewEncoder(w).Encode(book)
	})
	http.HandleFunc("/getAllBooks", func(w http.ResponseWriter, r *http.Request) {
		books := bookStore.GetAllBooks()
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByAuthor", func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		books := bookStore.GetAllBooksByAuthor(author)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByPrice", func(w http.ResponseWriter, r *http.Request) {
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByPrice(price)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitle", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		books := bookStore.GetAllBooksByTitle(title)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitleAndAuthor", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		author := r.URL.Query().Get("author")
		books := bookStore.GetAllBooksByTitleAndAuthor(title, author)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByTitleAndPrice", func(w http.ResponseWriter, r *http.Request) {
		title := r.URL.Query().Get("title")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByTitleAndPrice(title, price)
		json.NewEncoder(w).Encode(books)
	})
	http.HandleFunc("/getAllBooksByAuthorAndPrice", func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		price, _ := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
		books := bookStore.GetAllBooksByAuthorAndPrice(author, price)
		json.NewEncoder(w).Encode(books)
	})
	http.ListenAndServe(":8080", nil)
}
