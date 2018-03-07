package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("Pong!")
	fmt.Fprintf(w, "Pong!")
}

func main() {
	port := ":8080"

	if len(os.Getenv("PORT_LISTENER")) != 0 {
		port = os.Getenv("PORT_LISTENER")
		fmt.Printf("Using port %s...\n", port)
	}

	http.HandleFunc("/", HelloWorld)
	go log.Fatal(http.ListenAndServe(port, nil))

	for {
		_, err := http.Get("listener.svc.cluster.local:8080")
		fmt.Println("Ping!")
		if err != nil {
			fmt.Printf("Received error: %s\n", err.Error())
		}
		time.Sleep(5 * time.Second)
	}
}
