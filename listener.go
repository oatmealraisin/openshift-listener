package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	port := ":8080"

	if len(os.Getenv("PORT_LISTENER")) != 0 {
		port = os.Getenv("PORT_LISTENER")
		fmt.Printf("Using port %s...\n", port)
	}

	http.HandleFunc("/", HelloWorld)
	log.Fatal(http.ListenAndServe(port, nil))
}
