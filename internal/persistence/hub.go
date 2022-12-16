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
	_hubRepo        HubRepository
	loadHubRepoOnce sync.Once
)

type HubRepository interface {
	CreateHub(ctx context.Context, data model.Hub) (model.Hub, error)
	SearchHub(ctx context.Context, hubLocation string, teamType string) ([]model.Hub, error)
}

func Hub() HubRepository {
	if _hubRepo == nil {
		panic("persistence: hubRepo Repository not initiated")
	}

	return _hubRepo
}

func LoadHubGroupRepository(ctx context.Context) (err error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: false,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	loadHubRepoOnce.Do(func() {
		_hubRepo, err = gormI.NewHubGormRepoPSQL(ctx, db)
	})
	return
}

func MockHubRepo() {
	loadHubRepoOnce.Do(func() {
		_hubRepo = mock.NewMockHub()
	})
	return
}
