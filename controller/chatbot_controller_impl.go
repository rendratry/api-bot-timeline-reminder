package controller

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type ChatbotControllerImpl struct {
	ChatbotService service.ChatbotService
}

func NewChatbotController(chatboService service.ChatbotService) *ChatbotControllerImpl {
	return &ChatbotControllerImpl{
		ChatbotService: chatboService,
	}
}

func (controller *ChatbotControllerImpl) GetAllMessages(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	chatbotRequest := web.GetChatbotMessagesRequest{}
	helper.ReadFromRequestBody(request, &chatbotRequest)

	chatbotResponse := controller.ChatbotService.GetAllMessages(request.Context(), chatbotRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   chatbotResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
