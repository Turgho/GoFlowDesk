package ticket

import "time"

type TicketMessages struct {
	ID        string     `gorm:"primaryKey;type:char(36)" json:"id"`
	TicketID  string     `gorm:"type:char(36);not null" json:"ticket_id"`
	SenderID  string     `gorm:"type:char(36);not null" json:"sender_id"`
	Message   string     `gorm:"type:text;not null" json:"message"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
