package main

import (
	"fmt"
	"net/http"
)

func main() {
	// main server code
	http.HandleFunc("/books", booksShow)
	http.ListenAndServe(":5000", nil)
}

func booksShow(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello Viktor\n")
}
