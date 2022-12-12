package persistence

import (
	"context"
	"sync"
	"test/internal/model"
)

var (
	_payloadRepo        PayloadRepository
	loadPayloadRepoOnce sync.Once
)

type PayloadRepository interface {
	StorePayload(data model.Payload) error
}

func Payload() PayloadRepository {
	if _payloadRepo == nil {
		panic("persistence: payloadRepo Repository not initiated")
	}

	return _payloadRepo
}

func LoadPayloadRespositoryS3(ctx context.Context) (err error) {
	loadPayloadRepoOnce.Do(func() {
		_payloadRepo, err = newPayLoadRepoS3(ctx)
	})
	return
}

func LoadPayloadRespositoryMock(ctx context.Context) (err error) {
	loadPayloadRepoOnce.Do(func() {
		_payloadRepo, err = newPayLoadMock(ctx)
	})
	return
}
