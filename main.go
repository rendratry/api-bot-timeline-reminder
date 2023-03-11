package main

import (
	"api-bot-timeline-reminder/app"
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/middleware"
	"net/http"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    ":4000",
		Handler: authMiddleware,
	}
}

func main() {

	go app.ConsummerDelayRabbitMQ()
	go app.ConsummerSendEmail()
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

}
