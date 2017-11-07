package main

import (
	"log"
	"net/http"
)

func main() {
	router := NewRouter()
	server := http.ListenAndServe(":9090", router)
	log.Fatal(server)
}
