package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	. "gopkg.in/mgo.v2"
)

func Mongodb(s *Session) Adapter {
	return func(h httprouter.Handler) httprouter.Handler {
		return func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
}
