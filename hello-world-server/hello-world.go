package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port int
)

func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	fmt.Fprint(w, "Hello World!")
}

func init() {
	flag.IntVar(&port, "port", 3000, "HTTP Server Port")
	flag.Parse()
}

func main() {
	httpAddr := fmt.Sprintf(":%v", port)
	log.Printf("Listening to %v", httpAddr)

	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
