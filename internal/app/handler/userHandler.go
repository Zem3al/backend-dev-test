package handler

import (
	"encoding/json"
	"net/http"

	"dev/internal/service"
)

func UserCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.CreateUserCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	data, err := service.CreateUser(r.Context(), cmd)

	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}
