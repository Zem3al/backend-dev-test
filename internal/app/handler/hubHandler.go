package handler

import (
	"dev/internal/service"
	"encoding/json"
	"net/http"
)

func HubCreate(w http.ResponseWriter, r *http.Request) {
	cmd := service.CreateHubCommand{}

	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	data, err := service.CreateHub(r.Context(), cmd)
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}
