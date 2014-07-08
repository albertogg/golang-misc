package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/google/go-github/github"
	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 3000, "HTTP Server Port")
	flag.Parse()
}

func check(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func organizationHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL)

	params := mux.Vars(req)
	name := params["name"]

	client := github.NewClient(nil)

	orgs, _, err := client.Organizations.List(name, nil)
	check(w, err)

	jsonResp, err := json.Marshal(orgs)
	check(w, err)

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func main() {

	httpAddr := fmt.Sprintf(":%v", port)
	log.Printf("Listening to %v", httpAddr)

	router := mux.NewRouter()

	router.HandleFunc("/{name:[a-z]+}/orgs", organizationHandler)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
