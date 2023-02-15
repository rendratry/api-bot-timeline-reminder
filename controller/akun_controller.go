package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type AkunController interface {
	LoginAdmin(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	CreateAdmin(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	RegisterBot(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
