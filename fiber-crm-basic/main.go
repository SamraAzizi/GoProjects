package main

import (
	"fmt"

	"github.com/akhil/go-fiber-crm-basic/lead"

	"github.com/akhil/go-fiber-crm-basic/database"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.getLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.newLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "lead.db")
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connection opened to database")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("database migrated")
}
func main() {
	app := fiber.New()
	initDatabase()()
	setupRoutes(app)
	app.Listen(3000)
	defer database.DBConn.Close()

}
