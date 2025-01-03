package model

import (
	"time"

	"gorm.io/gorm"
)

type LoginRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
	Code  string `json:"code" binding:"required,len=4"` // For verification codes
}

type RegisterRequest struct {
	Email     string `json:"email" binding:"required,email"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name"`
}

type User struct {
	ID        int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Email     string         `json:"email" gorm:"index:,unique,composite:emaildeletedat" binding:"required,email"`
	FirstName string         `json:"first_name" gorm:"not null;size:255" binding:"required"`
	LastName  string         `json:"last_name" gorm:"size:255"`
	LastSeen  time.Time      `json:"last_seen" gorm:"default:CURRENT_TIMESTAMP"`
	Status    UserStatus     `json:"status" gorm:"type:varchar(10);default:'offline'"`
	CreatedAt time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at,omitempty" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index:,unique,composite:emaildeletedat" json:"-"`
}

type UserStatus string

const (
	Online  UserStatus = "online"
	Offline UserStatus = "offline"
)
