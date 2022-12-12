package model

import "time"

type User struct {
	ID        string     `gorm:"primaryKey;size:50"`
	Name      string     `gorm:"size:100;not null"`
	Age       int        `gorm:"not null"`
	TeamID    string     `gorm:"size:50;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (User) TableName() string {
	return "users"
}
