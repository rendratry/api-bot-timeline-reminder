// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"api-bot-timeline-reminder/app"
	"api-bot-timeline-reminder/controller"
	"api-bot-timeline-reminder/middleware"
	"api-bot-timeline-reminder/repository"
	"api-bot-timeline-reminder/service"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"net/http"
)

// Injectors from injector.go:

func InitializedServer() *http.Server {
	akunRepositoryImpl := repository.NewAkunRepositoryImpl()
	db := app.GetConnection()
	validate := validator.New()
	akunServiceImpl := service.NewUserServiceImpl(akunRepositoryImpl, db, validate)
	akunControllerImpl := controller.NewAkunControllerImpl(akunServiceImpl)
	otpRepositoryImpl := repository.NewOtpRepositoryImpl()
	otpServiceImpl := service.NewOtpServiceImpl(otpRepositoryImpl, db, validate)
	otpControllerImpl := controller.NewOtpControllerImpl(otpServiceImpl)
	router := app.NewRouter(akunControllerImpl, otpControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := NewServer(authMiddleware)
	return server
}

// injector.go:

var akunSet = wire.NewSet(repository.NewAkunRepositoryImpl, wire.Bind(new(repository.AkunRepository), new(*repository.AkunRepositoryImpl)), service.NewUserServiceImpl, wire.Bind(new(service.AkunService), new(*service.AkunServiceImpl)), controller.NewAkunControllerImpl, wire.Bind(new(controller.AkunController), new(*controller.AkunControllerImpl)))

var otpSet = wire.NewSet(repository.NewOtpRepositoryImpl, wire.Bind(new(repository.OtpRepository), new(*repository.OtpRepositoryImpl)), service.NewOtpServiceImpl, wire.Bind(new(service.OtpService), new(*service.OtpServiceImpl)), controller.NewOtpControllerImpl, wire.Bind(new(controller.OtpController), new(*controller.OtpControllerImpl)))
