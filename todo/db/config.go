package db

import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)


func DBConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB, models []&Interface{}) {
	for _, model := range models {
		db.AutoMigrate(model)
	}
}