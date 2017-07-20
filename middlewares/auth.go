package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func JWTAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	}
}
