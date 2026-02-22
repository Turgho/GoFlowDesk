package ticket

import "time"

type Ticket struct {
	ID          string `gorm:"primaryKey;type:char(36)" json:"id"`
	Title       string `gorm:"size:200;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Status      string `gorm:"size:50;not null;default:'open'" json:"status"`     // e.g., "open", "pending", "resolved", "closed"
	Priority    string `gorm:"size:50;not null;default:'medium'" json:"priority"` // e.g., "low", "medium", "high"
	RequesterID string `gorm:"type:char(36);not null" json:"requester_id"`
	AssigneeID  string `gorm:"type:char(36)" json:"assignee_id"`

	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"-"`
}
