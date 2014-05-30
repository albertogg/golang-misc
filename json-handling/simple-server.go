package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

var port int

func init() {
	flag.IntVar(&port, "port", 3000, "HTTP Server Port")
	flag.Parse()
}

type Bleh struct {
	Localtime time.Time `json:"localtime"`
	Hostname  string    `json:"hostname"`
}

type BlehJSON struct {
	Bleh Bleh `json:bleh"`
}

func blehHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("%s %s\n", req.Method, req.URL)
	w.Header().Set("Content-Type", "application/json")

	var blehJSON BlehJSON
	resp := blehJSON.Bleh
	resp.Localtime = time.Now()
	resp.Hostname, _ = os.Hostname()

	respData, err := json.Marshal(BlehJSON{Bleh: resp})
	if err != nil {
		panic(err)
	}

	w.Write(respData)
}

func main() {
	httpAddr := fmt.Sprintf(":%v", port)
	log.Printf("Listening to %v", httpAddr)

	router := mux.NewRouter()

	router.HandleFunc("/api/bleh", blehHandler).Methods("GET")
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe(httpAddr, nil))
}
