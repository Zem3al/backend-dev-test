package persistence

import (
	"context"
	"database/sql"
	"dev/internal/model"
)

type teamGorm struct {
	connection *sql.DB
}

func newTeamGormRepoPSQL(ctx context.Context, connection *sql.DB) (repo *teamGorm, err error) {
	return &teamGorm{connection: connection}, nil
}

func (t teamGorm) CreateTeam(ctx context.Context, data model.Team) (model.Team, error) {
	//TODO implement me
	panic("implement me")
}

func (t teamGorm) GetTeam(ctx context.Context) ([]model.Team, error) {
	//TODO implement me
	panic("implement me")
}
