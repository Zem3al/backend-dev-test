package mock

import (
	"context"

	"dev/internal/model"
)

type MockUser struct {
	users []model.User
}

func NewMockUser() *MockUser {
	return &MockUser{users: []model.User{}}
}

func (m *MockUser) CreateUser(ctx context.Context, data model.User) (model.User, error) {
	m.users = append(m.users, data)
	return data, nil
}
