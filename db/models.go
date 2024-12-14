package db

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        string `gorm:"primary_key;type:string;size:36"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	BaseModel `gorm:"embedded"`
	Email     string `gorm:"unique;not null;size:255"`
	Password  string `gorm:"not null"`
}

type Todo struct {
	BaseModel `gorm:"embedded"`
	Title     string    `gorm:"not null;size:100"`
	Expiry    time.Time `gorm:"not null"`
	Completed bool      `gorm:"not null;default:false"`
	UserID    string    `gorm:"type:string;not null;size:36"`
	User      User      `gorm:"foreignKey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
}

func (base *BaseModel) BeforeCreate(tx *gorm.DB) error {
	*base = BaseModel{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	return nil
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(&User{}, &Todo{})
}
