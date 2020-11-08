package main

import (
	"log"
	"net/http"
)

func main() {
	Router()
	log.Println("listen on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
