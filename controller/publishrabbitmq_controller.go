package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PublishRabbitMQController interface {
	PublishDelay(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
	SendEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params)
}
