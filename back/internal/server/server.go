package server

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	r "github.com/mkmuniz/FinTrack/internal/routes"
)

func Start() {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":3000"
	}

	app := fiber.New()
	
	r.SetupRoutes(app)
	
	log.Fatal(app.Listen(port))
}