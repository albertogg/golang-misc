package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 3000, "HTTP Server Port")
	flag.Parse()
}
func rootHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL)
	fmt.Fprint(w, "Hello World!")
}

func nameHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL)
	params := req.URL.Query()
	name := params.Get(":name")
	w.Write([]byte("Hello " + name))
}

func main() {
	httpAddr := fmt.Sprintf(":%v", port)
	log.Printf("Listening to %v", httpAddr)

	router := pat.New()

	router.Get("/", http.HandlerFunc(rootHandler))
	router.Get("/hello/:name", http.HandlerFunc(nameHandler))
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
