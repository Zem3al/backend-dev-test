package persistence

import (
	"context"
	"sync"

	"dev/internal/config"
	"dev/internal/model"
	gormI "dev/internal/persistence/gorm"
	"dev/internal/persistence/mock"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	_userRepo        UserRepository
	loadUserRepoOnce sync.Once
)

type UserRepository interface {
	CreateUser(ctx context.Context, data model.User) (model.User, error)
}

func User() UserRepository {
	if _userRepo == nil {
		panic("persistence: userRepo Repository not initiated")
	}

	return _userRepo
}

func LoadUserGroupRepository(ctx context.Context) (err error) {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.Get().PostgresURL,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	loadUserRepoOnce.Do(func() {
		_userRepo, err = gormI.NewUserGormRepoPSQL(ctx, db)
	})
	return
}

func MockUserRepo() {
	loadUserRepoOnce.Do(func() {
		_userRepo = mock.NewMockUser()
	})
	return
}
