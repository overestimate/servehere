package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	declareHandlers()
	fmt.Println("Listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
