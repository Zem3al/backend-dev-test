package model

import "time"

type Team struct {
	ID        string     `gorm:"primaryKey;size:50"`
	Name      string     `gorm:"size:100;not null"`
	Type      string     `gorm:"size:100;not null"`
	Users     []User     `gorm:"foreignKey:team_id"`
	HubID     string     `gorm:"size:50;not null"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}

func (Team) TableName() string {
	return "teams"
}
