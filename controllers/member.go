package controllers

import (
	// ""
	"net/http"
	"regexp"
	"time"

	"github.com/julienschmidt/httprouter"
	bcr "golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

type Member struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Email      string        `json:"email" bson:"email"`
	Name       string        `json:"name" bson:"name"`
	Password   string        `json:"password" bson:"password"`
	Created_at time.Time     `json:"created_at" bson:"created_at"`
}

const (
	DBName  = "purego"
	Cmember = "members"
)

var (
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

/*
	ENTRY POINT ----------------------------------------------------------------------------
*/

func (ct *Controller) EntryPoint(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	res := new(Response)
	res.Message = "Entry Point"

	ResponseAsJSON(w, res, http.StatusOK)
}

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
	return
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
	m.Name = r.Form.Get("name")
	m.Email = r.Form.Get("email")
	m.Created_at = time.Now()

	// VALIDATE NO EMPTY FORM VALUE
	if m.Name == "" || m.Email == "" || r.Form.Get("password") == "" {
		res.Errors = "Empty Value"
		ResponseError(w, res, http.StatusBadRequest)
		return
	}

	// VALIDATE EMAIL FORMAT
	if valid := emailRegexp.MatchString(m.Email); !valid {
		res.Errors = "Use a valid email format"
		ResponseError(w, res, http.StatusBadRequest)
		return
	}

	// FIND EMAIL AND VALIDATE UNIQUE EMAIL
	var emails []string
	if err := db.DB(DBName).C(Cmember).Find(nil).Distinct("email", &emails); err != nil {
		res.Errors = err.Error()
		ResponseError(w, res, http.StatusBadRequest)
		return
	}
	for _, email := range emails {
		if email == m.Email {
			res.Errors = "Email is already registered"
			ResponseError(w, res, http.StatusBadRequest)
			return
		}
	}

	hashedPw, err := bcr.GenerateFromPassword([]byte(r.Form.Get("password")), bcr.DefaultCost)
	if err != nil {
		res.Errors = err.Error()
		ResponseError(w, res, http.StatusBadRequest)
		return
	}
	m.Password = string(hashedPw)

	if err := db.DB(DBName).C(Cmember).Insert(&m); err != nil {
		res.Errors = err.Error()
		ResponseError(w, res, http.StatusBadRequest)
		return
	}

	res.Data = m
	res.Message = "OK"
	ResponseAsJSON(w, res, http.StatusCreated)
	return
}

/*
	LOGIN ONE USER ----------------------------------------------------------------------------
*/
