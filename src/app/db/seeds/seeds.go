package seeds

import (
	"app/db/scheme"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func CreateSample() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Create
	db.Create(&scheme.Talk{Msg: "サンプル"})
}
