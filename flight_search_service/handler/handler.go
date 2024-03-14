package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/repository"
	"github.com/olmandaniel/flight-tickets-sale/flight_search_service/services"
)

type FlightSearchHandler struct {
	repository.FlightSearchRepository
	*lib.ResponseService
}

func NewFlightSearchHandler(FlightSearchRepository repository.FlightSearchRepository, ResponseService *lib.ResponseService) *FlightSearchHandler {
	return &FlightSearchHandler{FlightSearchRepository: FlightSearchRepository, ResponseService: ResponseService}
}

func (hand *FlightSearchHandler) SearchFlightByParams(c *fiber.Ctx) error {
	origin := c.Query("origin")
	destination := c.Query("destination")
	date := c.Query("date")
	usecase := services.NewSearchFlightByParams(hand.FlightSearchRepository)
	flights, err := usecase.Execute(origin, destination, date)

	if err != nil {
		hand.ResponseService.SendError(c, 400, err.Error())
		return nil
	}
	hand.ResponseService.SendSuccess(c, http.StatusOK, flights)
	return nil
}

func (hand *FlightSearchHandler) GetFlightById(c *fiber.Ctx) error {
	id := c.Params("id")
	usecase := services.NewGetFlightById(hand.FlightSearchRepository)
	flight, err := usecase.Execute(id)

	if err != nil {
		hand.ResponseService.SendError(c, 400, err.Error())
		return nil
	}

	hand.ResponseService.SendSuccess(c, http.StatusOK, flight)
	return nil
}
