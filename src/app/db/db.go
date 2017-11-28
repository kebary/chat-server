package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	Database = connect()
)

func connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
