package utils

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type apiResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func ApiResponse(msg string) (response apiResponse) {
	return apiResponse{
		Success: true,
		Message: msg,
	}
}

type errResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error"`
}

func ErrResponse(err error, errMsg string) (respnse errResponse) {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		return errResponse{
			Success: false,
			Message: errMsg,
			Error:   ValidatorForError(err),
		}
	}

	return errResponse{
		Success: false,
		Message: errMsg,
		Error:   err.Error(),
	}
}
