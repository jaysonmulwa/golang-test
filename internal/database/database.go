package database

import (
	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

const (
	DB_TYPE = "sqlite3"
	DB_NAME = "./sample.db"
)

func Connect() (*gorm.DB, error) {

	db, err := gorm.Open(sqlite.Open(DB_NAME), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}
