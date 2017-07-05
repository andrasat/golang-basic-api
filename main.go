package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)

type Article struct {
  ID int `json:"id"`
  Title string `json:"title"`
  Desc string `json:"desc"`
  Content string `json:"content"`
}

type Articles []Article

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is a Homepage")
  fmt.Println("Endpoint: Homepage")
  fmt.Println("-----")
}

func getAllArticles(w http.ResponseWriter, r *http.Request) {
  articles := Articles{
    Article{Title: "test", Desc: "test", Content: "test"},
    Article{Title: "test", Desc: "test", Content: "test"},
  }
  fmt.Println("Endpoint: getAllArticles")

  json.NewEncoder(w).Encode(articles)
}

func getOneArticle(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  key := vars["id"]

  fmt.Fprintf(w, "Key: "+key)
}

func handleRequest() {
  myRouter := mux.NewRouter().StrictSlash(true)

  myRouter.HandleFunc("/", homePage)
  myRouter.HandleFunc("/all", getAllArticles)
  myRouter.HandleFunc("/article/{id}", getOneArticle)

  log.Fatal(http.ListenAndServe(":3000", myRouter))
}

func main() {
  fmt.Println("Rest API - Mux Router")
  handleRequest()
}