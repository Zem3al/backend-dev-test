package service

import (
	"context"

	"github.com/google/uuid"

	"dev/internal/model"
	"dev/internal/persistence"
)

func CreateHub(ctx context.Context, hubCommand CreateHubCommand) (hub CreateHubResponse, err error) {
	err = hubCommand.Validation()

	if err != nil {
		return hub, ErrInvalidInput
	}

	responseHub, err := persistence.Hub().CreateHub(ctx, model.Hub{
		ID:       uuid.New().String(),
		Name:     hubCommand.Name,
		Location: hubCommand.Location,
	})

	if err != nil {
		return hub, ErrDatabase
	}

	return CreateHubResponse{
		HubId:    responseHub.ID,
		Name:     responseHub.Name,
		Location: responseHub.Location,
	}, nil
}
