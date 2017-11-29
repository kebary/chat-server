package seeds

import (
	"app/db"
	"app/db/scheme"
)

func CreateSample() {
	// Create
	db.Database.Create(&scheme.Talk{Msg: "サンプル"})
}
