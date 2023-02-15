package controller

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"time"
)

type AkunControllerImpl struct {
	AkunService service.AkunService
}

func NewAkunControllerImpl(akunService service.AkunService) *AkunControllerImpl {
	return &AkunControllerImpl{
		AkunService: akunService,
	}
}

func (controller *AkunControllerImpl) LoginAdmin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginAdminRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.AkunService.LoginAdmin(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	jwt := helper.CreateNewJWT("bot.timeline.reminder", loginResponse.IdUser, "admin-bot-timeline", time.Hour*80)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
}

func (controller *AkunControllerImpl) CreateAdmin(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	akunCreateRequest := web.CreateAdminRequest{}
	helper.ReadFromRequestBody(request, &akunCreateRequest)

	akunResponse := controller.AkunService.CreateAdmin(request.Context(), akunCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   akunResponse,
	}

	jwt := helper.CreateNewJWT("bot.timeline.reminder", akunResponse.IdUser, "admin-bot-timeline", time.Hour*80)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
}

func (controller *AkunControllerImpl) RegisterBot(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	registerBotRequest := web.RegisterBotRequest{}
	helper.ReadFromRequestBody(request, &registerBotRequest)

	registerBotResponse := controller.AkunService.RegisterBot(request.Context(), registerBotRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   registerBotResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AkunControllerImpl) LoginUser(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	loginRequest := web.LoginUserRequest{}
	helper.ReadFromRequestBody(request, &loginRequest)

	loginResponse := controller.AkunService.LoginUser(request.Context(), loginRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   loginResponse,
	}

	jwt := helper.CreateNewJWT("bot.timeline.reminder", loginResponse.IdUser, "admin-bot-timeline", time.Hour*80)
	helper.WriteToResponseBodyWithJWT(writer, webResponse, jwt)
}

func (controller *AkunControllerImpl) InsertDataMahasiswa(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	insertDataRequest := web.InsertDataMahasiswaRequest{}
	helper.ReadFromRequestBody(request, &insertDataRequest)

	insertDataResponse := controller.AkunService.InsertDataMahasiswa(request.Context(), insertDataRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   insertDataResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *AkunControllerImpl) InsertDataStaff(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	insertDataRequest := web.InsertDataStaffRequest{}
	helper.ReadFromRequestBody(request, &insertDataRequest)

	insertDataResponse := controller.AkunService.InsertDataStaff(request.Context(), insertDataRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   insertDataResponse,
	}
	helper.WriteToResponseBody(writer, webResponse)
}
