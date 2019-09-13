package handler

import (
	"log"
	"net/http"

	"CASystem/interfaces/di"
	"CASystem/interfaces/response"
)

type demoGetResponse struct {
	Message string `json:"message"`
}

// HandleDemoGet is GET Demo message
func HandleDemoGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")

		msg, err := di.Demo.GetMessage()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		// response
		response.Success(
			writer,
			demoGetResponse{
				Message: msg,
			},
		)
	}
}
