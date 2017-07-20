package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ro "github.com/andrasat/pure-golang/routes"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

const (
	addr     = ":8080"
	serverDB = "localhost"
)

func main() {
	session, err := mgo.Dial(serverDB)
	if err != nil {
		log.Fatal("Mongo server error", err)
	}
	defer session.Close()

	r := httprouter.New()
	ro.Routes(r, session)

	server := &http.Server{
		Addr:           addr,
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("\n ------ TEST SERVER ------ \n")
	log.Println("Server running in " + addr)
	log.Fatal(server.ListenAndServe())
}
