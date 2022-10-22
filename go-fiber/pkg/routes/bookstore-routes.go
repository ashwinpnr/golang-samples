package routes

import (
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/controllers"
	"github.com/gofiber/fiber/v2"
)

func RegisterBookStoreRoutes(app *fiber.App) {
	app.Get("/api/v1/books", controllers.GetBooks)
	app.Get("/api/v1//book/:id", controllers.GetBookByID)
	app.Post("/api/v1//book", controllers.CreateBook)
	app.Put("/api/v1//book/:id", controllers.UpdateBook)
	app.Delete("/api/v1//book/:id", controllers.DeleteBook)
}
