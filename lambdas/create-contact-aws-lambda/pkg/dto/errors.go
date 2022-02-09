package dto

import (
	"errors"
	"fmt"
)

const (
	ValidationError = iota
	MalformedInputError
	InternalServerError
)

var (
	WrongRequestError = errors.New("invalid request")
	InvalidInputError = errors.New("malformed input")
	InsertionError    = errors.New("cannot insert a new item")
)

type LambdaError struct {
	ErrorCode  int `json:"error_code"`
	HttpStatus int `json:"http_status"`
	Err        error
}

func (e *LambdaError) Error() string {
	return fmt.Sprintf("code: %d, status: %d, msg: %s", e.ErrorCode, e.HttpStatus, e.Err.Error())
}
