package main

import (
	"context"
	"log"
	"os"
	"time"

	domain "github.com/Turgho/GoFlowDesk/internal/domain/user"
	repoUser "github.com/Turgho/GoFlowDesk/internal/repository/user"
	"github.com/Turgho/GoFlowDesk/internal/service/security"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx := context.Background()

	// Conexão com o banco
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL não definido")
	}
	pool, err := pgxpool.New(ctx, databaseURL)
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}
	defer pool.Close()

	// Repositórios
	userRepo := repoUser.NewUserRepository(pool)
	// ticketRepo := postgres.NewTicketRepository(pool) // desativado por enquanto

	// Hasher
	hasher := security.NewArgon2Hasher()

	// Seed de usuários
	if err := seedUsers(ctx, userRepo, hasher); err != nil {
		log.Println("Erro durante seed de usuários:", err)
	}

	// Seed de tickets (desativado)
	/*
		if err := seedTickets(ctx, ticketRepo, userRepo); err != nil {
			log.Println("Erro durante seed de tickets:", err)
		}
	*/

	log.Println("Seed finalizada!")
}

func seedUsers(ctx context.Context, userRepo *repoUser.UserRepository, hasher security.PasswordHasher) error {
	log.Println("Populando usuários...")

	users := []struct {
		Name  string
		Email string
		Role  domain.UserRole
		Pass  string
	}{
		{"Admin", "admin@goflowdesk.com", domain.UserRoleAdmin, "admin123"},
		{"Victor", "victor@goflowdesk.com", domain.UserRoleUser, "victor123"},
	}

	for _, u := range users {
		hash, err := hasher.HashPassword(u.Pass)
		if err != nil {
			log.Println("Erro ao gerar hash para:", u.Email, "-", err)
			continue
		}

		user := &domain.User{
			ID:           uuid.Must(uuid.NewV7()),
			Name:         u.Name,
			Email:        u.Email,
			PasswordHash: hash,
			Role:         u.Role,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		if err := userRepo.Create(ctx, user); err != nil {
			log.Println("Erro ao criar usuário:", u.Email, "-", err)
		} else {
			log.Println("Usuário criado:", u.Email)
		}
	}

	return nil
}

/*
func seedTickets(ctx context.Context, ticketRepo *postgres.TicketRepository, userRepo *postgres.UserRepository) error {
	log.Println("Populando tickets...")

	tickets := []struct {
		Title       string
		Description string
		Status      domain.TicketStatus
		OwnerEmail  string
	}{
		{"Teste 1", "Descrição do ticket 1", domain.TicketStatusOpen, "victor@goflowdesk.com"},
		{"Teste 2", "Descrição do ticket 2", domain.TicketStatusOpen, "victor@goflowdesk.com"},
	}

	for _, t := range tickets {
		owner, _ := userRepo.FindByEmail(ctx, t.OwnerEmail)
		if owner == nil {
			log.Println("Owner não encontrado para ticket:", t.Title)
			continue
		}

		ticket := &domain.Ticket{
			ID:          domain.NewUUID(),
			Title:       t.Title,
			Description: t.Description,
			Status:      t.Status,
			OwnerID:     owner.ID,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := ticketRepo.Save(ctx, ticket); err != nil {
			log.Println("Erro ao criar ticket:", t.Title, "-", err)
		} else {
			log.Println("Ticket criado:", t.Title)
		}
	}

	return nil
}
*/
