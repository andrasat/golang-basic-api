package routes

import (
	ctr "github.com/andrasat/pure-golang/controllers"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func Routes(r *httprouter.Router, session *mgo.Session) {
	c := &ctr.Controller{S: session}

	r.GET("/user/1", c.GetOneUser)
}
