package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	ro "github.com/andrasat/pure-golang/routes"
	"gopkg.in/mgo.v2"
)

const (
	addr     = ":3000"
	serverDB = "localhost"
)

func main() {
	session, err := mgo.Dial(serverDB)
	if err != nil {
		log.Fatal("Mongo server error", err)
	}
	defer session.Close()

	server := &http.Server{
		Addr:           addr,
		Handler:        ro.Routes(session),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Println("\n ------ TEST SERVER ------ \n")
	log.Println("Server running in " + addr)
	log.Fatal(server.ListenAndServe())
}
