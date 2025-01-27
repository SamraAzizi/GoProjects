package main

import (
	"fmt"

	"github.com/akhil/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get(getLeads)
	app.Get(GetLead)
	app.Post(newLead)
	app.Delete(DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "lead.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
}
func main() {
	app := fiber.New()
	initDatabase()()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
