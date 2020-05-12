package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port int
var status int

func handler(w http.ResponseWriter, r *http.Request) {
  log.Printf("-> :%d", port)
  log.Printf("<- %d", status)

	w.WriteHeader(status)
	io.WriteString(w, "")
}

func init() {
	flag.IntVar(&port, "port", 8080, "port on which to listen")
	flag.IntVar(&status, "status", 500, "status to return")
}

func main() {
	flag.Parse()
	http.HandleFunc("/", handler)

	// flag.PrintDefaults()
  fmt.Printf("⤾ deadendpoint is listening on port %d and will return %d\n", port, status)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
