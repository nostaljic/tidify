package models

import (
	"time"
)

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel
	UserEmail string `json:"user_email" gorm:"primary_key;varchar(50);not null;check:user_email <> ''"`
	SnsType   string `json:"sns_type" gorm:"not null;check:sns_type <> ''"`
}
