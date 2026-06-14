package models

import (
	"errors"
	"fmt"
)

type KotError struct {
	ReqId      string `json:"reqId"`
	ErrorCode  int    `json:"errorCode"`
	ErrorMsg   string `json:"errorMsg"`
	HttpStatus int    `json:"httpStatus"`
	Info       ReqInfo
}

func (e KotError) Error() string {
	return fmt.Sprintf("client error; reqId: '%s' errorCode: '%d' errorMsg: '%s'", e.ReqId, e.ErrorCode, e.ErrorMsg)
}

var (
	ErrBadRequestParams  = errors.New("bad request params")
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrNotFound          = errors.New("not found")
	ErrPayloadTooLarge   = errors.New("payload too large")
	ErrUnprocessable     = errors.New("unprocessable entity")
	ErrRateLimitExceeded = errors.New("rate limit exceeded")
	ErrInternalServer    = errors.New("internal server error")
)
