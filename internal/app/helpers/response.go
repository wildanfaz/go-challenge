package helpers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Err     string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewResponse(c *fiber.Ctx, statusCode int, message string, err error, data interface{}) error {
	if err != nil {
		return c.Status(statusCode).JSON(Response{
			Status:  checkStatus(statusCode),
			Message: message,
			Err:     err.Error(),
		})
	}

	return c.Status(statusCode).JSON(Response{
		Status:  checkStatus(statusCode),
		Message: message,
		Data:    data,
	})
}

func checkStatus(statusCode int) string {
	statusMessage := map[int]string{
		http.StatusInternalServerError: "Internal Server Error",
		http.StatusNotFound:            "Not Found",
		http.StatusBadRequest:          "Bad Request",
		http.StatusUnauthorized:        "Unauthorized",
		http.StatusUnprocessableEntity: "Unprocessable Entity",
		http.StatusOK:                  "OK",
		http.StatusCreated:             "Created",
	}

	msg, ok := statusMessage[statusCode]

	if !ok {
		return "Unknown Status Code"
	}

	return msg
}
