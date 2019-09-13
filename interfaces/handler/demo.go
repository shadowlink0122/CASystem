package handler

import (
	"log"
	"net/http"

	"github.com/shadowlink0122/CASystem/interfaces/di"
)

type demoGetResponse struct {
	Message string `json:"message"`
}

// HandleDemoGet is GET Demo message
func HandleDemoGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
				Message: msg
			}
		)
	}
}
