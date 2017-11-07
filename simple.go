package main

import (
	"fmt"
	"log"
	"net/http"
)

func getHash() string {
	return "1234"
}

func handler(w http.ResponseWriter, r *http.Request) {
	hash := getHash()
	fmt.Fprintf(w, "{\"hash\": %s}", hash)
	log.Println("Request")
}
