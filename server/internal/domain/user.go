package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id" gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	FirstName      string    `json:"first_name" gorm:"type:varchar(50);not null"`
	LastName       string    `json:"last_name" gorm:"type:varchar(50);not null"`
	Email          string    `json:"email" gorm:"type:text;not null;uniqueIndex"`
	HashPassword   *string   `json:"-"`
	AvatarURL      string    `json:"avatar_url" gorm:"type:text;default:''"`
	AvatarPublicID *string   `json:"-" `
	AuthProvider   string    `json:"provider" gorm:"type:text;not null;default:'local'"`
	AuthProviderID *string   `json:"-"`
	UserRole       string    `json:"role" gorm:"varchar(30);not null"`
	CreatedAt      time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt      time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
