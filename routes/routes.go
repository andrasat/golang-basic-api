package routes

import (
	ctr "github.com/andrasat/pure-golang/controllers"
	md "github.com/andrasat/pure-golang/middlewares"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func Routes(r *httprouter.Router, session *mgo.Session) {
	c := &ctr.Controller{S: session}

	r.GET("/member/:id", md.JWTAuth(c.GetOneUser))
	r.POST("/member", c.Register)
}
