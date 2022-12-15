package service

import (
	"context"

	"github.com/google/uuid"

	"dev/internal/model"
	"dev/internal/persistence"
)

func CreateUser(ctx context.Context, userCommand CreateUserCommand) (user CreateUserResponse, err error) {
	err = userCommand.Validation()

	if err != nil {
		return user, ErrInvalidInput
	}

	responseUser, err := persistence.User().CreateUser(ctx, model.User{
		ID:     uuid.New().String(),
		Name:   userCommand.Name,
		Age:    userCommand.Age,
		TeamID: userCommand.TeamID,
	})

	if err != nil {
		return user, ErrDatabase
	}

	return CreateUserResponse{
		UserID: responseUser.ID,
		Name:   responseUser.Name,
		Age:    responseUser.Age,
		TeamID: responseUser.TeamID,
	}, nil
}
