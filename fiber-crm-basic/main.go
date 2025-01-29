package main

import (
	"fmt"
	"log"

	"github.com/akhil/go-fiber-crm-basic/database"
	"github.com/akhil/go-fiber-crm-basic/lead"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func main() {
	app := fiber.New()

	// Initialize database
	database.InitDatabase()
	fmt.Println("Database connected and migrated")

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}
