package utils

type WebResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

func ApiResponse(msg string, err error) (response WebResponse) {
	var successTag bool
	switch err == nil {
	case true:
		successTag = true
	default:
		successTag = false
	}

	return WebResponse{
		Success: successTag,
		Message: msg,
		Error:   err,
	}
}
