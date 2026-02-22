package user

import "time"

type User struct {
	ID           string     `gorm:"primaryKey;type:char(36)" json:"id"`
	Name         string     `gorm:"size:100;not null" json:"name"`
	Email        string     `gorm:"size:150;unique;not null" json:"email"`
	PasswordHash string     `gorm:"size:255;not null" json:"-"`
	Role         string     `gorm:"size:50;not null;default:'user'" json:"role"` // e.g., "admin", "agent", "user"
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    *time.Time `gorm:"index" json:"-"`
}
