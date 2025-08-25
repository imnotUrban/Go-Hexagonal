package domain

import (
	"errors"
	"fmt"
)

const (
	ErrorCodeDuplicateKey        = "duplicate_key_error"
	ErrorCodeNotFound            = "not_found"
	ErrorCodeInternalServerError = "internal_server_error"
)

var ErrDuplicateKey = errors.New("duplicate key error")

type AppError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}
