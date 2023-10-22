package main

import (
	"github.com/gofiber/fiber/v2"
)

type response struct {
	Success bool
	Status  int
	Message string
	Payload any
}

func ApiResponses(err error, code int, message string, data any) *response {
	resultResponse := response{}

	if err != nil {
		resultResponse = response{
			Success: false,
			Status:  code,
			Message: fiber.ErrBadRequest.Message,
			Payload: data,
		}
		return &resultResponse
	}

	resultResponse = response{
		Success: true,
		Status:  code,
		Message: message,
		Payload: data,
	}
	return &resultResponse
}

type LogTracking struct {
	Message string
	Method  string
	Uri     string
	TraceId any
}
