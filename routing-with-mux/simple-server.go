package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var port int

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
	params := mux.Vars(req)
	name := params["name"]
	w.Write([]byte("Hello " + strings.Title(name)))
}

func main() {
	httpAddr := fmt.Sprintf(":%v", port)
	log.Printf("Listening to %v", httpAddr)

	router := mux.NewRouter()

	router.HandleFunc("/", rootHandler)
	router.HandleFunc("/hello/{name:[a-z]+}", nameHandler).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
