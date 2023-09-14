package main

import (
	"fmt"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome Gohelper\n")
}

func main() {
	http.ListenAndServeTLS(":8080", http.HandlerFunc(greeting))
}