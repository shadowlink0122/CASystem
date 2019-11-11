package handler

import (
	"log"
	"net/http"

	"CASystem/pkg/di"
	"CASystem/pkg/server/response"
)

type GetUserResponse struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
}

func UserGet() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		users, err := di.User.GetList()
		if err != nil {
			log.Println(err)
			response.InternalServerError(writer, "Internal Server Error")
			return
		}

		resp := usersToResponseBody(users)

		response.Success(
			writer,
			resp,
		)
	}
}

func usersToResponseBody(names []string) GetUserResponse {
	users := make([]User, 0, len(names))
	for i, _ := range names {
		user := User{Name: names[i]}
		users = append(users, user)
	}

	return GetUserResponse{Users: users}
}
