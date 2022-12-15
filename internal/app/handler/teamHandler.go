package handler

import (
	"dev/internal/service"
	"encoding/json"
	"net/http"
)

func TeamCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.CreateTeamCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	data, err := service.CreateTeam(r.Context(), cmd)
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}
