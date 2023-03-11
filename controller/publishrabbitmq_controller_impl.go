package controller

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PublishRabbitMQControllerImpl struct {
	PublishRabbitMQ service.PublishRabbitMQService
}

func NewPublishRabbitMQControllerImpl(publishRabbitMQ service.PublishRabbitMQService) *PublishRabbitMQControllerImpl {
	return &PublishRabbitMQControllerImpl{
		PublishRabbitMQ: publishRabbitMQ,
	}
}

func (controller *PublishRabbitMQControllerImpl) PublishDelay(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	publishRequest := web.PublishDelayRequest{FromApp: request.Header.Get("App-auth")}
	helper.ReadFromRequestBody(request, &publishRequest)

	publishResponse := controller.PublishRabbitMQ.PublishDelay(publishRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   publishResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *PublishRabbitMQControllerImpl) SendEmail(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	publishRequest := web.SendEmailRequest{FromApp: request.Header.Get("App-auth")}
	helper.ReadFromRequestBody(request, &publishRequest)

	publishResponse := controller.PublishRabbitMQ.SendEmail(publishRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   publishResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
