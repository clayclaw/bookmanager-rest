package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// For the Gigantic Teabagger's Fake IRL Library.

type Book struct {
	Name      string `json:"Name"`
	Issued    int    `json:"Issued"`
	Available bool   `json:"Available"`
	Author    string `json:"Author"`
	IssuedTo  string `json:"IssuedTo"`
}

var Books = []Book{
	Book{Name: "Why am I looking at stackoverflow", Issued: 122244, Available: true, Author: "Noob", IssuedTo: ""},
	Book{Name: "A brief history of time", Issued: 200, Available: true, Author: "Stephen William Hawking", IssuedTo: ""},
	Book{Name: "How to bake a cake", Issued: 7000, Available: true, Author: "Awesome Co0ker", IssuedTo: ""},
	Book{Name: "Chickens or eggs", Issued: 400, Available: true, Author: "PhilosopherX", IssuedTo: "CuriousKid"},
	Book{Name: "This is so hot", Issued: 100, Available: true, Author: "Hotland solider", IssuedTo: ""},
	Book{Name: "My Longest Yeah Boi Ever", Issued: 450, Available: true, Author: "Good boy", IssuedTo: ""},
	Book{Name: "Dice game", Issued: 9990, Available: true, Author: "17 years old kid", IssuedTo: ""},
	Book{Name: "Snake candy", Issued: 140, Available: true, Author: "Anna", IssuedTo: ""},
	Book{Name: "Jerk Chicken", Issued: 10014, Available: false, Author: "Unknown", IssuedTo: "Giant2"},
}

func BooksAvailable(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Endpoint hit: BookAvailable\n")

	var filter []Book

	for i, book := range Books {
		if book.Available {
			filter = append(filter, book)
			println(i, book.Name)
		}
	}

	json.NewEncoder(w).Encode(filter)

}

func BookStatus(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("Endpoint hit: BookStatus\n")

	json.NewEncoder(w).Encode(Books)
	
}

func MostIssued(w http.ResponseWriter, r *http.Request)  {

	fmt.Printf("Endpoint hit: MostIssued\n")

	var top = Books[0]

	for i, book := range Books {
		if top.Issued < book.Issued {
			top = book
			println(i, book.Name)
		}
	}

	json.NewEncoder(w).Encode(top)

}

func TopTrending(w http.ResponseWriter, r *http.Request){

	fmt.Printf("Endpoint hit: TopTrending\n")

	var trending = [3]Book {Books[0], Books[1], Books[2]}
	json.NewEncoder(w).Encode(trending)

}

func BookIssued(w http.ResponseWriter, r *http.Request){

	fmt.Printf("Endpoint hit: BookIssued\n")

	var issued []Book

	for i, book := range Books {
		if book.IssuedTo != "" {
			issued = append(issued, book)
			println(i, book.Name)
		}
	}

	json.NewEncoder(w).Encode(issued)

}

const port = "8080"

func main() {

	http.HandleFunc("/api/available", BooksAvailable)
	http.HandleFunc("/api/status", BookStatus)
	http.HandleFunc("/api/mostIssued", MostIssued)
	http.HandleFunc("/api/top", TopTrending)
	http.HandleFunc("/api/issued", BookIssued)

	fmt.Printf("Http server initialized\n")

	log.Fatal(http.ListenAndServe(":"+port, nil))

}
