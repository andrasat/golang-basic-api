package main_test

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ro "github.com/andrasat/pure-golang/routes"
	"gopkg.in/mgo.v2"
)

var (
	server  *httptest.Server
	reader  io.Reader
	userUrl string
	session *mgo.Session
)

const serverDB = "localhost"

func TestRegister(t *testing.T) {

	session, err := mgo.Dial(serverDB)
	if err != nil {
		fmt.Errorf("Mongo Error : %v", err)
	}
	defer session.Close()

	server = httptest.NewServer(ro.Routes(session))

	userUrl = fmt.Sprintf("%s/", server.URL)

	userForm := `{"name": "testing", "email": "test@test.com", "pass": "testing"}`

	reader = strings.NewReader(userForm)

	req, err := http.NewRequest("POST", userUrl+"member", reader)
	req.Header = map[string]string{
		"Content-Type": "application/x-www-form-urlencoded",
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != 201 {
		t.Errorf("\nStatus Code is : %d, \nExpected : 201", res.StatusCode)
	}

}
