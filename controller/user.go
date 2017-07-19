package controller

import (
	"net/http"

	"gopkg.in/mgo.v2"
)

type User struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Token    string `json:"token,omitempty"`
	Password string `json:"password,omitempty"`
}

func (ct *Controller) GetAllUser(w http.ResponseWriter, r *http.Request) {

	DB := ct.s.Copy()
	defer DB.Close()
	res := new(Response)

	ResponseAsJSON(w, res, http.StatusOK)
}
