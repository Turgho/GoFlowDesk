package database

import (
	"database/sql"

	"github.com/Turgho/GoFlowDesk/internal/domain/ticket"
	"github.com/Turgho/GoFlowDesk/internal/domain/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func SetupGorm(db *sql.DB) *gorm.DB {
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return gormDB
}

func AutoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&user.User{},
		&ticket.Ticket{},
		&ticket.TicketMessages{},
	)
	if err != nil {
		return err
	}
	return nil
}
