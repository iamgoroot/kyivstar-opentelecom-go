package models

import (
	"fmt"
)

type KotError struct {
	ReqId      string `json:"reqId"`
	ErrorCode  int    `json:"errorCode"`
	ErrorMsg   string `json:"errorMsg"`
	HttpStatus int    `json:"httpStatus"`
}

func (e KotError) Error() string {
	return fmt.Sprintf("client error; reqId: '%s' errorCode: '%d' errorMsg: '%s'", e.ReqId, e.ErrorCode, e.ErrorMsg)
}
