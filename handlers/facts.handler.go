package handlers

import (
	"fmt"
	"sort"
	"time"

	"github.com/emarifer/go-htmx-trivia/database"
	"github.com/emarifer/go-htmx-trivia/models"
	"github.com/gofiber/fiber/v2"
)

// This handler is no longer in use
/* func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Enrique from Go Trivia App!!")
} */

var year = time.Now().Year()

func ListFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Sp.DB.From("facts").Select("*").Execute(&facts)

	sort.SliceStable(facts, func(i, j int) bool {
		return facts[i].ID > facts[j].ID
	}) // VER nota abajo.

	// return c.Status(fiber.StatusOK).JSON(facts) // for API REST
	return c.Render("index", fiber.Map{
		"PageTitle": "Gopher Trivia Time",
		"Title":     "Gopher Trivia Time",
		"Subtitle":  "Facts for funtimes with friends!",
		"Facts":     facts,
		"Year":      year,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"PageTitle": "Create New Fact",
		"Title":     "New Fact",
		"Subtitle":  "Add a cool fact!",
		"Year":      year,
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := models.Fact{
		Question: c.FormValue("question"),
		Answer:   c.FormValue("answer"),
	}

	// Validamos que los campos vengan y sean correctos
	errors := models.ValidateStruct(fact)
	if errors != nil {
		c.Status(fiber.StatusBadRequest).SendString("The 'question' and/or 'answer' fields must be completed")
		return nil
	}

	facts := []models.Fact{}
	err := database.DB.Sp.DB.From("facts").Insert(fact).Execute(&facts)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Something went wrong…")

		return nil
	}

	c.Set("HX-Redirect", "/")

	return nil
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Select("*").Single().Eq("id", id).Execute(&fact)
	if err != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"PageTitle": fmt.Sprintf("Show Fact #%s", id),
		"Title":     "Single Fact",
		"Fact":      fact,
		"Year":      year,
	})
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Select("*").Single().Eq("id", id).Execute(&fact)
	if err != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"PageTitle": fmt.Sprintf("Edit Fact #%s", id),
		"Title":     "Edit Fact",
		"Subtitle":  fmt.Sprintf("Edit your interesting fact #%s", id),
		"Fact":      fact,
		"Year":      year,
	})
}

func UpadteFact(c *fiber.Ctx) error {
	id := c.Params("id")

	fact := models.Fact{
		Question: c.FormValue("question"),
		Answer:   c.FormValue("answer"),
	}

	// Validamos que los campos vengan y sean correctos
	errors := models.ValidateStruct(fact)
	if errors != nil {
		c.Status(fiber.StatusBadRequest).SendString("The 'question' and/or 'answer' fields must be completed")
		return nil
	}

	// Solución la problema de error con el método Update. VER:
	// https://github.com/nedpals/supabase-go/issues/3#issuecomment-1542984127
	var results []models.Fact
	err := database.DB.Sp.DB.From("facts").Update(fact).Eq("id", id).Execute(&results)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString("Something went wrong…")

		return nil
	}

	c.Set("HX-Redirect", fmt.Sprintf("/fact/%s", id))

	return nil
}

func DeleteFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Delete().Eq("id", id).Execute(&facts)
	if err != nil {
		return NotFound(c)
	}

	c.Set("HX-Redirect", "/")

	return nil
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).Render("404", fiber.Map{
		"PageTitle": "404 | Gopher Trivia Time",
	})
}

// This handler is no longer in use
/* func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"PageTitle": "Confirmation Page",
		"Title":     "Fact added successfully!",
		"Subtitle":  "Add more wonderful facts to the list!",
	})
} */

/* DOCUMENTACIÓN supabase-go
https://github.com/nedpals/supabase-go
*/

/*
https://stackoverflow.com/questions/70618200/how-to-implement-a-redirect-with-htmx
https://www.reddit.com/r/htmx/comments/ot6kai/comment/h6v5cn9/?utm_source=share&utm_medium=web2x&context=3
https://www.reddit.com/r/htmx/comments/s36zx2/how_do_you_use_hxredirect/
https://docs.gofiber.io/api/ctx/#set

https://strapengine.com/go-fiber-custom-header-middleware/

https://htmx.org/attributes/hx-confirm/
*/

/*
COMO EL DRIVER DE SUPABASE PARA GO, EN PRINCIPIO, NO DA ORDENACIÓN
DE LOS ELEMENTOS DE LA TABLA, HACEMOS UNA ORDENACIÓN MANUAL DEL SLICE:
https://freshman.tech/snippets/go/sorting-in-go/
*/

/* handle errors with HTMX:
   https://htmx.org/extensions/response-targets/
*/
