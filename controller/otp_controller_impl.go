package controller

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/web"
	"api-bot-timeline-reminder/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type OtpControllerImpl struct {
	OtpService service.OtpService
}

func NewOtpControllerImpl(otpService service.OtpService) *OtpControllerImpl {
	return &OtpControllerImpl{
		OtpService: otpService,
	}
}

func (Controller *OtpControllerImpl) SendOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	otpRequest := web.OtpRequest{}
	helper.ReadFromRequestBody(request, &otpRequest)

	otpResponse := Controller.OtpService.SendOtp(request.Context(), otpRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   otpResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (Controller *OtpControllerImpl) VerifikasiOtp(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	otpValidationRequest := web.OtpValidationRequest{}
	helper.ReadFromRequestBody(request, &otpValidationRequest)

	otpValidationResponse := Controller.OtpService.OtpValidation(request.Context(), otpValidationRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   otpValidationResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
