package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
