package main

import (
	"fmt"
	"log"
	"net/http"

	ro "github.com/andrasat/pure-golang/routes"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

const (
	addr = ":8080"
)

func main() {
	session, err := mgo.Dial("mongodb://localhost/purego")
	if err != nil {
		log.Fatal("Mongo server error")
	}

	r := httprouter.New()
	ro.Routes(r, session)

	fmt.Println("\n ------ TEST SERVER ------ \n")
	log.Println("Server running in " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
