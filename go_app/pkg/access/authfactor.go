package access

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Authentication factors

type AuthFactor struct {
	gorm.Model
	ID          uint `gorm:"type:integer;not null; autoincrement;primary_key"`
	Email       bool
	Phone       bool
	Token       bool
	CodeTemp    string
	TokenActive string
	UserID      uuid.UUID
	CreatedAt   *time.Time     `gorm:"not null;default:now"`
	UpdatedAt   *time.Time     `gorm:"not null;default:now"`
	DeletedAt   gorm.DeletedAt `gorm:"not null;index"`
}
