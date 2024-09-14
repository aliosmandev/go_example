package server

import (
	"github.com/gofiber/fiber/v2"

	user "go_starter/cmd/api/rest/routes"
	"go_starter/internal/database"
)

type FiberServer struct {
	*fiber.App

	db database.Service
}

func New() *FiberServer {
	server := &FiberServer{
		App: fiber.New(fiber.Config{
			ServerHeader: "go_starter",
			AppName:      "go_starter",
		}),

		db: database.New(),
	}

	user.Register(server.App)

	return server
}
