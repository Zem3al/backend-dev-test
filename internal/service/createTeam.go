package service

import (
	"context"

	"github.com/google/uuid"

	"dev/internal/model"
	"dev/internal/persistence"
)

func CreateTeam(ctx context.Context, teamCommand CreateTeamCommand) (team CreateTeamResponse, err error) {
	err = teamCommand.Validation()

	if err != nil {
		return team, ErrInvalidInput
	}

	responseTeam, err := persistence.Team().CreateTeam(ctx, model.Team{
		ID:    uuid.New().String(),
		Name:  teamCommand.Name,
		Type:  teamCommand.Type,
		HubID: teamCommand.HubID,
	})

	if err != nil {
		return team, ErrDatabase
	}

	return CreateTeamResponse{
		TeamID: responseTeam.ID,
		Name:   responseTeam.Name,
		Type:   responseTeam.Type,
		HubID:  responseTeam.HubID,
	}, nil
}
