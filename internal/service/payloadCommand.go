package service

import (
	"github.com/asaskevich/govalidator"
	"test/internal/model"
	"test/internal/persistence"
	"time"
)

type PayloadCommandStore struct {
	WriterKey string    `json:"writerKey" valid:"required"`
	TypeSDK   string    `json:"typeSDK"  valid:"required"`
	SentAt    time.Time `json:"sentAt" valid:"required"`
	Batch     []struct {
		TimeStamp   time.Time              `json:"timeStamp" valid:"required"`
		Name        string                 `json:"name" valid:"required"`
		RequestID   string                 `json:"requestId" valid:"required"`
		AnonymousID string                 `json:"anonymousId" valid:"required"`
		WriterKey   string                 `json:"writerKey" valid:"required"`
		Context     map[string]interface{} `json:"context"`
	} `json:"batch"`
}

func (cmd *PayloadCommandStore) Valid() error {
	_, err := govalidator.ValidateStruct(cmd)
	return err
}

func (cmd *PayloadCommandStore) Run() error {
	PayloadRepo := persistence.Payload()

	err := cmd.Valid()

	if err != nil {
		return err
	}

	pRepo := model.Payload{
		Batchs:    []*model.Batch{},
		WriterKey: &cmd.WriterKey,
		TypeSDK:   &cmd.TypeSDK,
		SentAt:    &cmd.SentAt,
	}

	for _, batch := range cmd.Batch {
		pRepo.Batchs = append(pRepo.Batchs, &model.Batch{
			TimeStamp:   &batch.TimeStamp,
			Name:        &batch.Name,
			RequestID:   &batch.RequestID,
			AnonymousID: &batch.AnonymousID,
			WriterKey:   &batch.WriterKey,
			Context:     batch.Context,
		})
	}

	return PayloadRepo.StorePayload(pRepo)
}
