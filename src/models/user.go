package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `json:"create_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	BaseModel
	UserEmail string `gorm:"primary_key;varchar(50);not null;" json:"user_email"`
	SnsType   string `gorm:"not null;" json:"sns_type"`
}
