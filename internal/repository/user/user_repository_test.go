package user_test

import (
	"context"
	"os"
	"testing"
	"time"

	domain "github.com/Turgho/GoFlowDesk/internal/domain/user"
	repo "github.com/Turgho/GoFlowDesk/internal/repository/user"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func setupTestRepo(t *testing.T) (*repo.UserRepository, context.Context, func()) {
	ctx := context.Background()
	databaseURL := os.Getenv("DATABASE_URL_TEST")
	if databaseURL == "" {
		t.Fatal("DATABASE_URL_TEST não definido")
	}
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		t.Fatal("Erro ao conectar no banco:", err)
	}

	// inicia transação para limpar ao final
	tx, err := pool.Begin(ctx)
	if err != nil {
		t.Fatal("Erro ao iniciar transação:", err)
	}

	cleanup := func() {
		_ = tx.Rollback(ctx)
		pool.Close()
	}

	repoInstance := repo.NewUserRepository(pool)
	return repoInstance, ctx, cleanup
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	repo, ctx, cleanup := setupTestRepo(t)
	t.Cleanup(cleanup)

	user := &domain.User{
		ID:           uuid.Must(uuid.NewV7()),
		Name:         "Test User",
		Email:        "testuser@example.com",
		PasswordHash: "hash",
		Role:         domain.UserRoleUser,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	if err := repo.Create(ctx, user); err != nil {
		t.Fatal("Erro ao criar usuário:", err)
	}

	found, err := repo.FindByID(ctx, user.ID)
	if err != nil {
		t.Fatal("Erro ao buscar usuário por ID:", err)
	}
	if found.Email != user.Email {
		t.Fatalf("Esperado %s, achado %s", user.Email, found.Email)
	}

	found2, err := repo.FindByEmail(ctx, user.Email)
	if err != nil {
		t.Fatal("Erro ao buscar usuário por email:", err)
	}
	if found2.ID != user.ID {
		t.Fatalf("Esperado ID %s, achado %s", user.ID, found2.ID)
	}

	users, err := repo.List(ctx, 10, 0)
	if err != nil {
		t.Fatal("Erro ao listar usuários:", err)
	}
	if len(users) == 0 {
		t.Fatal("Esperado pelo menos 1 usuário, encontrado 0")
	}
}
