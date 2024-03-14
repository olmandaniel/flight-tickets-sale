package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/olmandaniel/flight-tickets-sale/api_gateway/dtos"
	"github.com/olmandaniel/flight-tickets-sale/api_gateway/lib"
)

type ApiGatewayHandler struct {
	lib.Env
	*lib.ResponseService
}

func NewApiGatewayHandler(env lib.Env, responseService *lib.ResponseService) *ApiGatewayHandler {
	return &ApiGatewayHandler{
		Env:             env,
		ResponseService: responseService,
	}
}

func (gateway *ApiGatewayHandler) SearchFlight(c *fiber.Ctx) error {
	date := c.Query("date")
	origin := c.Query("origin")
	destination := c.Query("destination")
	url := fmt.Sprintf("%s/flights?date=%s&origin=%s&destination=%s", gateway.Env.SearchFlightService, date, origin, destination)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	if response.StatusCode != http.StatusOK {
		var errorResponse lib.ErrorResponse
		json.Unmarshal(body, &errorResponse)
		gateway.ResponseService.SendError(c, response.StatusCode, errorResponse.Error)
		return nil
	}

	var successResponse lib.SuccessResponse
	json.Unmarshal(body, &successResponse)
	gateway.ResponseService.SendSuccess(c, response.StatusCode, successResponse.Data)
	return nil
}

func (gateway *ApiGatewayHandler) FindFlight(c *fiber.Ctx) error {
	flightId := c.Params("id")
	url := fmt.Sprintf("%s/flights/%s", gateway.Env.SearchFlightService, flightId)
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	if response.StatusCode != http.StatusOK {
		var errorResponse lib.ErrorResponse
		json.Unmarshal(body, &errorResponse)
		gateway.ResponseService.SendError(c, response.StatusCode, errorResponse.Error)
		return nil
	}

	var successResponse lib.SuccessResponse
	json.Unmarshal(body, &successResponse)
	gateway.ResponseService.SendSuccess(c, response.StatusCode, successResponse.Data)
	return nil
}

func (gateway *ApiGatewayHandler) BookFlight(c *fiber.Ctx) error {

	bookFlightRequest := new(dtos.BookFlightRequest)

	if err := c.BodyParser(bookFlightRequest); err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	url := fmt.Sprintf("%s/book-flights", gateway.Env.BookingService)
	client := http.Client{}
	jsonData, err := json.Marshal(&bookFlightRequest)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	payload := strings.NewReader(string(jsonData))
	request, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	request.Header.Add("Content-Type", "application/json")
	response, err := client.Do(request)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		gateway.ResponseService.SendError(c, http.StatusInternalServerError, err.Error())
		return nil
	}
	if response.StatusCode != http.StatusCreated {
		var errorResponse lib.ErrorResponse
		json.Unmarshal(body, &errorResponse)
		gateway.ResponseService.SendError(c, response.StatusCode, errorResponse.Error)
		return nil
	}

	var successResponse lib.SuccessResponse
	json.Unmarshal(body, &successResponse)
	gateway.ResponseService.SendSuccess(c, response.StatusCode, successResponse.Data)
	return nil
}
