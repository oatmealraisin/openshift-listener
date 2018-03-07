package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func HelloWorld(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Pong!")
	fmt.Fprintf(w, "Pong!")
}

func main() {
	host := "listener.svc.cluster.local"
	port := "8080"

	if len(os.Getenv("PORT_LISTENER")) != 0 {
		port = os.Getenv("PORT_LISTENER")
		fmt.Printf("Using port %s...\n", port)
	}

	if len(os.Getenv("HOST_LISTENER")) != 0 {
		host = os.Getenv("HOST_LISTENER")
		fmt.Printf("Using host %s...\n", host)
	}

	http.HandleFunc("/", HelloWorld)
	go func() {
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}()

	for {
		time.Sleep(5 * time.Second)
		_, err := http.Get(fmt.Sprintf("http://%s:%s", host, port))
		fmt.Println("Ping!")
		if err != nil {
			fmt.Printf("Received error: %s\n", err.Error())
		}
	}
}
