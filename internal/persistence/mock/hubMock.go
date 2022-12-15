package mock

import (
	"context"

	"dev/internal/model"
)

type MockHub struct {
	hubs []model.Hub
}

func NewMockHub() *MockHub {
	return &MockHub{hubs: []model.Hub{}}
}

func (m *MockHub) CreateHub(ctx context.Context, data model.Hub) (model.Hub, error) {
	m.hubs = append(m.hubs, data)
	return data, nil
}

func (m *MockHub) SearchHub(ctx context.Context, hubLocation string, teamType string) ([]model.Hub, error) {
	return m.hubs, nil
}
