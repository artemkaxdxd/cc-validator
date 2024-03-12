package config

import (
	"errors"
	"net/http"
)

var ( // errors
	ErrInvalidYear       = errors.New("invalid year")
	ErrInvalidMonth      = errors.New("invalid month")
	ErrInvalidCVV        = errors.New("invalid cvv")
	ErrExpiredCard       = errors.New("credit card expired")
	ErrInvalidCardNumber = errors.New("invalid credit card number")
)

type ServiceCode int

const (
	CodeOK ServiceCode = iota
	CodeBadRequest
	CodeInvalidCardNumber
	CodeExpiredCard
	CodeInvalidCVV
)

func ErrToServiceCode(err error) ServiceCode {
	switch err {
	case nil:
		return CodeOK
	case ErrInvalidCardNumber:
		return CodeInvalidCardNumber
	case ErrInvalidYear,
		ErrInvalidMonth,
		ErrExpiredCard:
		return CodeExpiredCard
	case ErrInvalidCVV:
		return CodeInvalidCVV
	default:
		return CodeBadRequest
	}
}

func CodeToHttpStatus(code ServiceCode) int {
	switch code {
	case CodeOK:
		return http.StatusOK
	case CodeBadRequest,
		CodeInvalidCardNumber,
		CodeExpiredCard,
		CodeInvalidCVV:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}
