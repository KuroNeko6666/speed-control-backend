package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        string   `json:"id" gorm:"primaryKey"`
	Name      string   `json:"name"`
	Username  string   `json:"username" gorm:"unique"`
	Email     string   `json:"email" gorm:"unique"`
	Password  string   `json:"-" gorm:"password"`
	Role      string   `json:"role"`
	Devices   []Device `json:"-" gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (data *User) BeforeCreate(tx *gorm.DB) (err error) {
	data.ID = uuid.NewString()
	return
}
