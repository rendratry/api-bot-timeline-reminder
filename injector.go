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

func InitializedServer() *http.Server {
	wire.Build(
		app.GetConnection,
		validator.New,
		akunSet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
