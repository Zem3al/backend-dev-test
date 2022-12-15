package mock

import (
	"context"

	"dev/internal/model"
)

type TeamHub struct {
	teams []model.Team
}

func NewMockTeam() *TeamHub {
	return &TeamHub{teams: []model.Team{}}
}

func (t *TeamHub) CreateTeam(ctx context.Context, data model.Team) (model.Team, error) {
	t.teams = append(t.teams, data)
	return data, nil
}

func (t *TeamHub) GetTeam(ctx context.Context) ([]model.Team, error) {
	//TODO implement me
	panic("implement me")
}
