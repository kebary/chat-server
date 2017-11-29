package migration

import (
	"app/db"
	"app/db/scheme"
)

func Migration() {
	// Migrate the schema
	db.Database.AutoMigrate(&scheme.Talk{})
}
