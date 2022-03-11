package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_TYPE = "sqlite3"
	DB_NAME = "./sample.db"
)

/*import (
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
  )*/

func Connect() (*sql.DB, error) {

	db, err := sql.Open(DB_TYPE, DB_NAME)
	if err != nil {
		return db, err
	}
	return db, nil
}
