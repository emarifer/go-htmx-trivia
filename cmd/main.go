package main

import (
	"github.com/emarifer/go-htmx-trivia/database"
	"github.com/emarifer/go-htmx-trivia/handlers"
	"github.com/emarifer/go-htmx-trivia/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	database.ConnectDB()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	routes.SetupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	app.Listen(":3000")
}

/* Reference:
https://divrhino.com/articles/rest-api-docker-go-fiber-from-scratch/
https://divrhino.com/articles/full-stack-go-fiber-with-docker-postgres/
https://divrhino.com/articles/crud-go-fiber-docker-postgres/

Inspired design:
https://github.com/marco-souza/gx
https://github.com/NerdCademyDev/gophat
*/
