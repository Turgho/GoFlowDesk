package user

import (
	"context"

	domain "github.com/Turgho/GoFlowDesk/internal/domain/user"
	repo "github.com/Turgho/GoFlowDesk/internal/repository/user"
	"github.com/Turgho/GoFlowDesk/internal/service/security"
)

// Service contains business logic related to users.
type Service struct {
	repo   *repo.UserRepository
	hasher security.PasswordHasher
}

// NewService constructs a user service.
func NewService(r *repo.UserRepository, hasher security.PasswordHasher) *Service {
	return &Service{repo: r, hasher: hasher}
}

// CreateUser performs validation, hashing and persistence.
func (s *Service) CreateUser(ctx context.Context, name, email, password string) (*domain.User, error) {
	// verify not already exists
	existing, err := s.repo.FindByEmail(ctx, email)
	if err != nil && err != domain.ErrUserNotFound {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrEmailAlreadyExists
	}

	hash, err := s.hasher.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user, err := domain.NewUser(name, email, hash, domain.UserRoleUser)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}
