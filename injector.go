//go:build wireinject
// +build wireinject

package main

import (
	"api-bot-timeline-reminder/app"
	"api-bot-timeline-reminder/controller"
	"api-bot-timeline-reminder/middleware"
	"api-bot-timeline-reminder/repository"
	"api-bot-timeline-reminder/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var akunSet = wire.NewSet(
	repository.NewAkunRepositoryImpl,
	wire.Bind(new(repository.AkunRepository), new(*repository.AkunRepositoryImpl)),
	service.NewUserServiceImpl,
	wire.Bind(new(service.AkunService), new(*service.AkunServiceImpl)),
	controller.NewAkunControllerImpl,
	wire.Bind(new(controller.AkunController), new(*controller.AkunControllerImpl)),
)

var otpSet = wire.NewSet(
	repository.NewOtpRepositoryImpl,
	wire.Bind(new(repository.OtpRepository), new(*repository.OtpRepositoryImpl)),
	service.NewOtpServiceImpl,
	wire.Bind(new(service.OtpService), new(*service.OtpServiceImpl)),
	controller.NewOtpControllerImpl,
	wire.Bind(new(controller.OtpController), new(*controller.OtpControllerImpl)),
)

var chatbotSet = wire.NewSet(
	repository.NewChatbotRepositoryImpl,
	wire.Bind(new(repository.ChatbotRepository), new(*repository.ChatbotRepositoryImpl)),
	service.NewChatbotService,
	wire.Bind(new(service.ChatbotService), new(*service.ChatbotServiceImpl)),
	controller.NewChatbotController,
	wire.Bind(new(controller.ChatbotController), new(*controller.ChatbotControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		chatbotSet,
		akunSet,
		otpSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
