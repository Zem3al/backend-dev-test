package persistence

import (
	"context"
	"encoding/json"
	"log"
	"test/internal/model"
	"time"
)

type PayLoadMock struct {
}

func (p PayLoadMock) StorePayload(data model.Payload) error {
	_, err := json.Marshal(data)

	if err != nil {
		return err
	}
	time.Sleep(1 * time.Second)
	log.Println("Uploaded data")
	return nil
}

func newPayLoadMock(ctx context.Context) (repo *PayLoadMock, err error) {
	return &PayLoadMock{}, nil
}
