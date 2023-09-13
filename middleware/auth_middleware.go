package middleware

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(Handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{Handler: Handler}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if "RAHASIA" == request.Header.Get("X-API-Key") {
		// Ok
		middleware.Handler.ServeHTTP(writer, request)
	} else {
		// Error
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)

		WebResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteToResponseBody(writer, WebResponse)
	}
}
