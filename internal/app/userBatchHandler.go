package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"test/internal/service"
)

func UploadData(w http.ResponseWriter, r *http.Request) {
	cmd := service.PayloadCommandStore{}
	r.Body = http.MaxBytesReader(w, r.Body, MAX_BYTES_READ)
	err := json.NewDecoder(r.Body).Decode(&cmd)
	if err != nil {
		ResponeError(w, ErrUnProcessableEnity)
		return
	}

	workerService, err := service.GetWorkerService()

	if err != nil {
		fmt.Println(err)
		ResponeError(w, ErrBadRequest)
		return
	}

	err = workerService.AddToQueue(&cmd)

	if err != nil {
		fmt.Println(err)
		ResponeError(w, ErrBadRequest)
		return
	}

	ResponeData(w, nil)
}
