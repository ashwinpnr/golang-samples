package main

import (
	"log"
	"strconv"

	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/config"
	"github.com/ashwinpnr/golang-samples/go-fiber/pkg/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()
	routes.RegisterBookStoreRoutes(app)
	log.Printf("Starting Server")
	serveraddr := ":" + strconv.Itoa(config.GetConfig().Server.Port)
	if err := app.Listen(serveraddr); err != nil {
		log.Fatal(err)
	}
}
