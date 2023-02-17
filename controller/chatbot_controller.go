package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ChatbotController interface {
	GetAllMessages(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
