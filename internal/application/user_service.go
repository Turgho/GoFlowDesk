package application

import (
	"log/slog"

	"github.com/Turgho/GoFlowDesk/internal/domain/user"
	"github.com/Turgho/GoFlowDesk/internal/infrastructure/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(name, email, password string) (*user.User, error)
	FindUserByID(id string) (*user.User, error)
	FindUserByEmail(email string) (*user.User, error)
}

type userService struct {
	db       *gorm.DB
	userRepo user.UserRepository
	logger   *slog.Logger
}

func NewUserService(userRepo user.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// ------------------------------

// CreateUser implements [UserService].
func (s *userService) CreateUser(name, email, password string) (*user.User, error) {
	logger := s.logger.With("method", "CreateUser", "email", email)

	// Generate a new UUID
	id, err := uuid.NewV7()
	if err != nil {
		logger.Error("Failed to generate UUID", "error", err)
		return nil, err
	}

	// Create user entity
	newUser := &user.User{
		ID:           id.String(),
		Name:         name,
		Email:        email,
		PasswordHash: password,
		Role:         user.RoleUser,
	}

	// Transaction
	if err := s.db.Transaction(func(tx *gorm.DB) error {
		return s.userRepo.Create(tx, newUser)
	}); err != nil {
		if database.IsUniqueViolation(err) {
			logger.Info("Attempt to create user with duplicate email")
			return nil, user.ErrEmailAlreadyExists
		}
		logger.Error("Failed to create user", "error", err)
		return nil, err
	}

	logger.Info("User created successfully", "user_id", newUser.ID)
	return newUser, nil
}

// FindUserByEmail implements [UserService].
func (s *userService) FindUserByEmail(email string) (*user.User, error) {
	panic("unimplemented")
}

// FindUserByID implements [UserService].
func (s *userService) FindUserByID(id string) (*user.User, error) {
	panic("unimplemented")
}
