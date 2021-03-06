package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s.\n", r.URL.Path[1:])
}

func main() {
	http.ListenAndServe(":8000", http.HandlerFunc(handler))
}
