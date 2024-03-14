package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/dtos"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/lib/msgqueue"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/repository"
	"github.com/olmandaniel/flight-tickets-sale/booking_service/services"
)

type BookingHandler struct {
	Emitter         msgqueue.EventEmitter
	ResponseService *lib.ResponseService
	repository.BookingRepository
}

func NewBookingHandler(Emitter msgqueue.EventEmitter, ResponseService *lib.ResponseService, BookingRepository repository.BookingRepository) *BookingHandler {
	return &BookingHandler{Emitter: Emitter, ResponseService: ResponseService, BookingRepository: BookingRepository}
}

func (hand *BookingHandler) BookFlight(c *fiber.Ctx) error {
	bookFlightRequest := new(dtos.BookFlightRequest)

	if err := c.BodyParser(bookFlightRequest); err != nil {
		hand.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	usecase := services.NewBookFlight(hand.BookingRepository, hand.Emitter)
	err := usecase.Execute(*bookFlightRequest)

	if err != nil {
		hand.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	hand.ResponseService.SendSuccess(c, http.StatusCreated, nil)
	return nil
}
