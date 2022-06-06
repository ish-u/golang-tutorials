package main

import (
	"fmt"
	"log"
	"net/http"
)

// /hello route handler
// w - Response
// r - Request
func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "METHOD is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err %v\n", err)
		return
	}

	fmt.Fprintf(w, "POST Request Successful\n")
	name := r.FormValue("name")
	quote := r.FormValue("quote")
	fmt.Fprintf(w, "Name = %s \n Quote = %s\n", name, quote)
}

func main() {
	// Makes Go Lang look for the index.html inside the static directory
	fileServer := http.FileServer(http.Dir("./static"))

	// handle functions to handle HTTP Requests on the routes
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server Starting at PORT 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
