package middlewares

import (
	"log"
	"net/http"
	"strings"

	ctr "github.com/andrasat/pure-golang/controllers"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

func JWTAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		res := new(ctr.Response)

		tokenString := strings.Split(r.Header.Get("Authorization"), " ")
		log.Printf("\nString splitted: %v", tokenString[1])

		token, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
			return []byte("SECRET"), nil
		})
		if err != nil {
			log.Printf("Error: %v", err)
			res.Message = "JWT Parse Error !"
			ctr.ResponseError(w, res, http.StatusInternalServerError)
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Printf("Username: %v", claims["username"])
			h(w, r, ps)
			return
		}

		res.Message = "You are unauthorized"
		ctr.ResponseError(w, res, http.StatusUnauthorized)
		return
	}
}
