package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Watch this space!</h1>")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	fmt.Println("Starting server on port 80...")
	http.ListenAndServe(":80", nil)
}
