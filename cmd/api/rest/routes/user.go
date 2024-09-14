package user

import (
	"go_starter/internal/usecase/user"

	"github.com/gofiber/fiber/v2"
)

func Register(app *fiber.App) {
	group := app.Group("/users")
	group.Get("/", user.GetUser)
}
