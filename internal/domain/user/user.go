package user

import (
	"time"

	"github.com/gofrs/uuid"
)

// User represents an account in the system.  Most business logic lives
// in the service layer; the entity itself is intentionally "dumb".
type User struct {
	ID           uuid.UUID  `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	PasswordHash string     `json:"-"`
	Role         UserRole   `json:"role"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at,omitempty"`
}

type UserRole string

const (
	UserRoleAdmin UserRole = "admin"
	UserRoleUser  UserRole = "user"
	UserRoleGuest UserRole = "agent"
)

func (r UserRole) IsValid() bool {
	switch r {
	case UserRoleAdmin, UserRoleUser, UserRoleGuest:
		return true
	default:
		return false
	}
}

// NewUser constructs and validates a new User entity.
func NewUser(name, email, passwordHash string, role UserRole) (*User, error) {
	if name == "" {
		return nil, ErrEmptyName
	}
	if email == "" {
		return nil, ErrEmptyEmail
	}
	if passwordHash == "" {
		return nil, ErrEmptyPassword
	}
	if !role.IsValid() {
		return nil, ErrInvalidUserRole
	}

	now := time.Now().UTC()
	return &User{
		ID:           uuid.Must(uuid.NewV7()),
		Name:         name,
		Email:        email,
		PasswordHash: passwordHash,
		Role:         role,
		CreatedAt:    now,
		UpdatedAt:    now,
	}, nil
}
