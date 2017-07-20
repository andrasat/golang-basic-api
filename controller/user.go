package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (ct *Controller) GetOneUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	DB := ct.S.Copy()
	defer DB.Close()
	res := new(Response)

	res.Message = "Test Message"
	res.Data = "This is the data"

	ResponseAsJSON(w, res, http.StatusOK)
}
