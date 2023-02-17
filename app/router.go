package app

import (
	"api-bot-timeline-reminder/controller"
	"api-bot-timeline-reminder/exception"
	"github.com/julienschmidt/httprouter"
)

func NewRouter(
	AkunController controller.AkunController,
	OtpController controller.OtpController,
	ChatbotController controller.ChatbotController,
) *httprouter.Router {
	router := httprouter.New()

	router.POST("/bot/api/adminregister", AkunController.CreateAdmin)
	router.POST("/bot/api/adminlogin", AkunController.LoginAdmin)
	router.POST("/bot/api/registerbot", AkunController.RegisterBot)
	router.POST("/bot/api/userlogin", AkunController.LoginUser)
	router.POST("/bot/api/insertdatamahasiswa", AkunController.InsertDataMahasiswa)
	router.POST("/bot/api/insertdatastaff", AkunController.InsertDataStaff)
	router.POST("/bot/api/sendotp", OtpController.SendOtp)
	router.POST("/bot/api/otpvalidation", OtpController.VerifikasiOtp)
	router.POST("/bot/api/getchatbotmessages", ChatbotController.GetAllMessages)

	router.PanicHandler = exception.ErrorHandler

	return router

}
