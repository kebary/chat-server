package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	Echo = echo.New()
)

func init() {
	// db
	Migration()
	createSample()
}

func createSample() {
	db, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()
	// Create
	db.Create(&Talk{Msg: "サンプル"})

	// Read
	// var product Product
	// db.First(&product, 1)
	// db.First(&product, "code = ?", "L1212")

	// Update - update product's price to 2000
	// db.Model(&product).Update("Price", 2000)

	// Delete - delete product
	// db.Delete(&product)
}

func main() {
	// Middleware
	Echo.Use(middleware.Logger())
	Echo.Use(middleware.Recover())
	// routes
	routes()
	// start
	Echo.Logger.Fatal(Echo.Start(":1323"))
}
