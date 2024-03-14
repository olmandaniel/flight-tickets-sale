package lib

import "github.com/gofiber/fiber/v2"

type ErrorResponse struct {
	ErrorCode int    `json:"error_code"`
	Error     string `json:"error"`
}

type SuccessResponse struct {
	Data any `json:"data"`
}

type ResponseService struct {
}

func NewResponseService() *ResponseService {
	return &ResponseService{}
}

func (*ResponseService) SendError(c *fiber.Ctx, status int, errorMessage string) {
	c.Status(status).JSON(ErrorResponse{ErrorCode: status, Error: errorMessage})
}

func (*ResponseService) SendSuccess(c *fiber.Ctx, status int, data any) {
	c.Status(status).JSON(SuccessResponse{Data: data})
}
