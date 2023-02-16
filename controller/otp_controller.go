package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OtpController interface {
	SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	VerifikasiOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
