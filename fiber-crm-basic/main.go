package main

import (
	"github.com/gofiber/fiber"
)

func setupRoutes(app *fiber.App) {
	app.Get(getLeads)
	app.Post(newLead)
	app.Delete(DeleteLead)
}
func main() {
	app := fiber.New()
	setupRoutes(app)

}
