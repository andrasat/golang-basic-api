package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	addr = ":8080"
)

func main() {
	r := httprouter.New()

	log.Println("Server running in " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
