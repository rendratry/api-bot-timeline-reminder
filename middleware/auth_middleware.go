package middleware

import (
	"api-bot-timeline-reminder/helper"
	"api-bot-timeline-reminder/model/domain"
	"api-bot-timeline-reminder/model/web"
	"log"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	XapiKey := "12345678"
	if request.Header.Get("Authorization") != "" {
		jwt := request.Header.Get("Authorization")
		iss := domain.JwtClaims{}
		iss = helper.ValidateJWT(jwt, "admin-bot-timeline")
		status := helper.GetIssuer(iss.Subject)
		log.Println(iss.Subject)
		if status && request.Header.Get("X-Api-Key") == XapiKey {
			// ok
			middleware.Handler.ServeHTTP(writer, request)
		} else {
			// error
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusUnauthorized)

			webResponse := web.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: "UNAUTHORIZED",
			}

			helper.WriteToResponseBody(writer, webResponse)
		}
	} else {
		if request.URL.Path == "/bot/api/adminlogin" || request.URL.Path == "/bot/api/adminregister" ||
			request.URL.Path == "/bot/api/registerbot" || request.URL.Path == "/bot/api/userlogin" ||
			request.URL.Path == "/bot/api/sendotp" || request.URL.Path == "/bot/api/otpvalidation" {
			if request.Header.Get("X-Api-Key") == XapiKey {
				// ok
				middleware.Handler.ServeHTTP(writer, request)
			} else {
				// error
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}

				helper.WriteToResponseBody(writer, webResponse)
			}
		} else {
			if request.URL.Path == "/bot/api/publishdelay" || request.URL.Path == "/bot/api/sendemail" {
				authapp := helper.GetRegisteredApp(request.Header.Get("App-auth"))
				if request.Header.Get("X-Api-Key") == XapiKey && authapp {
					// ok
					middleware.Handler.ServeHTTP(writer, request)
				} else {
					// error
					writer.Header().Set("Content-Type", "application/json")
					writer.WriteHeader(http.StatusUnauthorized)

					webResponse := web.WebResponse{
						Code:   http.StatusUnauthorized,
						Status: "UNAUTHORIZED",
					}

					helper.WriteToResponseBody(writer, webResponse)
				}
			} else {
				// error
				writer.Header().Set("Content-Type", "application/json")
				writer.WriteHeader(http.StatusUnauthorized)

				webResponse := web.WebResponse{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
				}

				helper.WriteToResponseBody(writer, webResponse)
			}

		}
	}

}
