package model

import (
	"time"
)

type Device struct {
	ID        string       `json:"id" gorm:"primaryKey;not null"`
	UserID    string       `json:"user_id" gorm:"size:191"`
	Model     string       `json:"model"`
	Address   string       `json:"address"`
	User      User         `json:"user" gorm:"foreignKey:UserID"`
	Data      []DeviceData `json:"data" gorm:"foreignKey:DeviceID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
