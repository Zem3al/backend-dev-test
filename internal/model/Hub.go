package model

import "time"

type Hub struct {
	ID        string     `gorm:"primaryKey;size:50"`
	Name      string     `gorm:"size:100;not null"`
	Location  string     `gorm:"size:100;not null"`
	Teams     []Team     `gorm:"foreignKey:hub_id"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Hub) TableName() string {
	return "hubs"
}
