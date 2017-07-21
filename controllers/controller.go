package controllers

import (
	"encoding/json"
	. "gopkg.in/mgo.v2"
	"net/http"
)

type Controller struct {
	S *Session
}

type Response struct {
	Errors  string      `json:"error,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseAsJSON(w http.ResponseWriter, r *Response, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("WWW-Authenticate", "Basic-realm=TEST SERVER")
	w.WriteHeader(status)

	result, _ := json.Marshal(r)
	w.Write(result)
}

func ResponseError(w http.ResponseWriter, r *Response, status int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("WWW-Authenticate", "Basic-realm=TEST SERVER")
	w.WriteHeader(status)

	result, _ := json.Marshal(r)
	w.Write(result)
}
