package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	gorm.Model
	BaseModel `gorm:"embedded"`
	Email     string `gorm:"unique;not null;length:255"`
	Password  string `gorm:"not null"`
}

type Todo struct {
	gorm.Model
	BaseModel `gorm:"embedded"`
	Title     string    `gorm:"not null;length:255"`
	Expiry    time.Time `gorm:"not null"`
	Completed bool      `gorm:"not null;default:false"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

func Models() []interface{} {
	return []interface{}{
		&User{},
		&Todo{},
	}
}
