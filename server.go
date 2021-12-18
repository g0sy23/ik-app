package ik_app

import (
	"time"

	"github.com/g0sy23/ik-app/internal/handler"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Server struct {
	App *fiber.App
}

func NewServer(handler *ik_handler.Handler) *Server {
	s := &Server{
		App: fiber.New(fiber.Config{
			AppName: 				"ik-app",
			ReadTimeout: 		10 * time.Second,
			WriteTimeout:		10 * time.Second,
		}),
	}
	s.App.Use(cors.New(cors.Config{
		AllowMethods:	"GET, POST, PATCH, DELETE",
	}))
	handler.InitRoutes(s.App)
	return s
}

func (s *Server) Run(port string) error {
	return s.App.Listen(":" + port)
}

func (s *Server) Shutdown() error {
	return s.App.Shutdown()
}
