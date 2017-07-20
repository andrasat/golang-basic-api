package controllers

import (
	// "log"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	bcr "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Member struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Password   string        `json:"password" bson:"password"`
	Created_at time.Time     `json:"created_at" bson:"created_at"`
}

const (
	DBName  = "purego"
	Cmember = "members"
)

/*
	GET ONE USER ----------------------------------------------------------------------------
*/

func (ct *Controller) GetOneUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db := ct.S.Copy()
	defer db.Close()
	res := new(Response)

	res.Message = "Test Message"
	res.Data = "This is the data"

	ResponseAsJSON(w, res, http.StatusOK)
}

/*
	REGISTER ONE USER ----------------------------------------------------------------------------
*/

func (ct *Controller) Register(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	db := ct.S.Copy()
	defer db.Close()
	res := new(Response)
	var m Member

	r.ParseForm()
	m.ID = bson.NewObjectId()
	m.Name = strings.Join(r.Form["name"], "")
	m.Created_at = time.Now()

	hashedPw, err := bcr.GenerateFromPassword([]byte(strings.Join(r.Form["password"], "")), bcr.DefaultCost)
	if err != nil {
		res.Errors = err
		ResponseError(w, res, http.StatusBadRequest)
		return
	}
	m.Password = string(hashedPw)

	if err := db.DB(DBName).C(Cmember).Insert(&m); err != nil {
		res.Errors = err
		ResponseError(w, res, http.StatusBadRequest)
		return
	}

	res.Data = m
	res.Message = "OK"
	ResponseAsJSON(w, res, http.StatusOK)
	return
}

/*
	LOGIN ONE USER ----------------------------------------------------------------------------
*/
