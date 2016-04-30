package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

t

func fooHandler(w http.ResponseWriter, r *http.Request) {

}

// Get entries from database.
// Serve as JSON to be consumed by browser.

func main() {

	http.Handle("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
