package ik_server

import (
	"time"

	"github.com/g0sy23/ik-app/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	app *fiber.App
}

func New(handler *ik_handler.Handler) *Server {
	s := &Server{
		app: fiber.New(fiber.Config{
			AppName:      "ik-app",
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		}),
	}
	s.app.Use(cors.New(cors.Config{
		AllowMethods: "GET, POST, PATCH, DELETE",
	}))
	handler.InitRoutes(s.app)
	s.app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).SendString("Sorry can't find that!")
	})
	return s
}

func (s *Server) Run(port string) error {
	return s.app.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	return s.app.Shutdown()
}
