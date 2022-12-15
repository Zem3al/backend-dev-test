package gorm

import (
	"context"

	"gorm.io/gorm"

	"dev/internal/model"
)

type teamGorm struct {
	connection *gorm.DB
}

func NewTeamGormRepoPSQL(ctx context.Context, connection *gorm.DB) (repo *teamGorm, err error) {
	return &teamGorm{connection: connection}, nil
}

func (t teamGorm) CreateTeam(ctx context.Context, data model.Team) (model.Team, error) {
	result := t.connection.Create(&data).Model(model.Team{})

	if result.Error != nil {
		return model.Team{}, result.Error
	}

	return data, nil
}

func (t teamGorm) GetTeam(ctx context.Context) ([]model.Team, error) {
	//TODO implement me
	panic("implement me")
}
