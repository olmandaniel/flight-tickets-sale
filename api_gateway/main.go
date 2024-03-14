package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/api_gateway/handler"
	"github.com/olmandaniel/flight-tickets-sale/api_gateway/lib"
)

func main() {
	app := fiber.New()

	env := lib.NewEnv()
	responseService := lib.NewResponseService()
	gatewayHandler := handler.NewApiGatewayHandler(env, responseService)
	apiRoute := app.Group("/api/v1")
	r := apiRoute.Group("/flights")
	r.Get("/", gatewayHandler.SearchFlight)
	r.Get("/:id", gatewayHandler.FindFlight)
	r.Post("/book", gatewayHandler.BookFlight)

	log.Fatal(app.Listen(fmt.Sprintf(":%s", env.AppPort)))
}
