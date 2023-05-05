package routes

import (
	"github.com/gofiber/fiber/v2"
	"srm_arch/internal/controller/http/v1"
)

func RegisterRoutes(app *fiber.App, c *v1.Controller) {
	routes := app.Group("/register")
	routes.Post("/login/", c.Login)
}
