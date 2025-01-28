package lead

import (
	"github.com/akhil/go-fiber-crm-basic/database"
	"github.com/jinzhu/gorm"
)

type Lead struct {
	gorm.Model
	Name    string
	Company string
	Email   string
	Phone   int
}

func GetLeads(c *fiber.Ctx) {
	db := database.DBConn
	var leads []lead
	db.Find(&leads)
	c.JSON(leads)

}

func GetLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead lead
	db.JSON(lead)
}

func NewLead(c *fiber.Ctx) {

}

func DeleteLead(c *fiber.Ctx) {

}
