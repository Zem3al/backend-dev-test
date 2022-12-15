package gorm

import (
	"context"

	"gorm.io/gorm"

	"dev/internal/model"
)

type userGorm struct {
	connection *gorm.DB
}

func NewUserGormRepoPSQL(ctx context.Context, connection *gorm.DB) (repo *userGorm, err error) {
	return &userGorm{connection: connection}, nil
}

func (u userGorm) CreateUser(ctx context.Context, data model.User) (model.User, error) {
	result := u.connection.Create(&data).Model(model.User{})
	if result.Error != nil {
		return model.User{}, result.Error
	}

	return data, nil
}
