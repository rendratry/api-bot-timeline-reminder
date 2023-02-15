package main

import (
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

	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
