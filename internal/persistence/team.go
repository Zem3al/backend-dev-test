package persistence

import (
	"context"
	"dev/internal/config"
	"dev/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	_teamRepo        TeamRepository
	loadTeamRepoOnce sync.Once
)

type TeamRepository interface {
	CreateTeam(ctx context.Context, data model.Team) (model.Team, error)
	GetTeam(ctx context.Context) ([]model.Team, error)
}

func Team() TeamRepository {
	if _payloadRepo == nil {
		panic("persistence: teamRepo Repository not initiated")
	}

	return _teamRepo
}

func LoadTeamGroupRespository(ctx context.Context) (err error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	sqlDB, err := db.DB()

	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

	loadTeamRepoOnce.Do(func() {
		_teamRepo, err = newTeamGormRepoPSQL(ctx, sqlDB)
	})
	return
}
