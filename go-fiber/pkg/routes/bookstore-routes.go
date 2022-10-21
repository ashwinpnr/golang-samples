package routes

import (
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookStoreRoutes(app *fiber.App) {
	app.Get("/books", controllers.GetBooks)
	app.Get("/book/:id", controllers.GetBookByID)
	app.Post("/book", controllers.CreateBook)
	app.Put("/book/:id", controllers.UpdateBook)
	app.Delete("/book/:id", controllers.DeleteBook)
}
