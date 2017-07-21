package routes

import (
	ctr "github.com/andrasat/pure-golang/controllers"
	md "github.com/andrasat/pure-golang/middlewares"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func Routes(session *mgo.Session) *httprouter.Router {
	c := &ctr.Controller{S: session}
	r := httprouter.New()

	r.GET("/", c.EntryPoint)
	r.GET("/member/:id", md.JWTAuth(c.GetOneUser))
	r.POST("/member", c.Register)

	return r
}
