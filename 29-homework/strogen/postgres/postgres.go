package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
	db, _ := gorm.Open(postgres.Open("postgres://postgres:0509@localhost:5432/bobur?sslmode=disable"))
	return db, nil
}
