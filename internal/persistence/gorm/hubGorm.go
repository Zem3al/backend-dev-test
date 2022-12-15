package gorm

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"dev/internal/model"
)

type hubGorm struct {
	connection *gorm.DB
}

func NewHubGormRepoPSQL(ctx context.Context, connection *gorm.DB) (repo *hubGorm, err error) {
	return &hubGorm{connection: connection}, nil
}

func (h hubGorm) CreateHub(ctx context.Context, data model.Hub) (model.Hub, error) {
	result := h.connection.Create(&data).Model(model.Hub{})

	if result.Error != nil {
		return model.Hub{}, result.Error
	}

	return data, nil
}

func (h hubGorm) SearchHub(ctx context.Context, hubLocation string, teamType string) ([]model.Hub, error) {
	var hubs []model.Hub
	var teamArgs []interface{}
	var hubArgs []interface{}

	if hubLocation != "" {
		hubArgs = append(hubArgs, "hubs.location LIKE ?", fmt.Sprintf("%%%s%%", hubLocation))
	}

	if teamType != "" {
		teamArgs = append(teamArgs, "teams.type LIKE ?", fmt.Sprintf("%%%s%%", teamType))
	}

	result := h.connection.Model(model.Hub{}).Preload("Teams", teamArgs...).Preload("Teams.Users").Find(&hubs, hubArgs...)

	if result.Error != nil {
		return nil, result.Error
	}

	return hubs, nil
}
