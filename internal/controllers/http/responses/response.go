package responses

import "solidgate-test/config"

type ValidationResponse struct {
	Valid bool  `json:"valid"`
	Error Error `json:"error,omitempty"`
}

type Error struct {
	Code    config.ServiceCode `json:"code"`
	Message string             `json:"message"`
}

func NewValidationErr(code config.ServiceCode, msg string) ValidationResponse {
	return ValidationResponse{
		Error: Error{
			Code:    code,
			Message: msg,
		},
	}
}

func NewValidationResponse(valid bool) ValidationResponse {
	return ValidationResponse{
		Valid: valid,
	}
}
