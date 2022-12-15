package handler

import (
	"net/http"

	"dev/internal/service"
)

func Search(w http.ResponseWriter, r *http.Request) {
	cmd := service.SearchCommand{}
	cmd.Type = r.URL.Query().Get("type")
	cmd.Location = r.URL.Query().Get("location")

	data, err := service.Search(r.Context(), cmd)
	if err != nil {
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, data)
}
