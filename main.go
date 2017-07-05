package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
)

type Article struct {
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

func handleRequest() {
  http.HandleFunc("/", homePage)
  http.HandleFunc("/all", getAllArticles)
  log.Fatal(http.ListenAndServe(":3000", nil))
}

func main() {
  handleRequest()
}