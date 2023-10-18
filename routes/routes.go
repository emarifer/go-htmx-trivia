package routes

import (
	"github.com/emarifer/go-htmx-trivia/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFact)

	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)

	app.Get("/fact/:id", handlers.ShowFact)

	app.Get("/fact/:id/edit", handlers.EditFact)
	app.Patch("/fact/:id", handlers.UpadteFact)

	app.Delete("/fact/:id", handlers.DeleteFact)
}
