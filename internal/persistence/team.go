package persistence

import (
	"context"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"dev/internal/config"
	"dev/internal/model"
	gormI "dev/internal/persistence/gorm"
	"dev/internal/persistence/mock"
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
	if _teamRepo == nil {
		panic("persistence: teamRepo Repository not initiated")
	}

	return _teamRepo
}

func LoadTeamGroupRepository(ctx context.Context) (err error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: false, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	loadTeamRepoOnce.Do(func() {
		_teamRepo, err = gormI.NewTeamGormRepoPSQL(ctx, db)
	})
	return
}

func MockTeamRepo() {
	loadTeamRepoOnce.Do(func() {
		_teamRepo = mock.NewMockTeam()
	})
	return
}
