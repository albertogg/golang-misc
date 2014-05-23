package main

import (
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Printf("Listening to 3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
