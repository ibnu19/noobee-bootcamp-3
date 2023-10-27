package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func ApiResponse[V any](msg string, payload V) response {
	response := response{
		Success: true,
		Message: msg,
		Payload: payload,
	}

	return response
}

type errResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrorResponse(msg string, errMsg error) errResponse {
	var ve validator.ValidationErrors
	if errors.As(errMsg, &ve) {
		return errResponse{
			Success: false,
			Message: msg,
			Error:   ValidatorErrors(errMsg),
		}
	}

	return errResponse{
		Success: false,
		Message: msg,
		Error:   errMsg.Error(),
	}

}
