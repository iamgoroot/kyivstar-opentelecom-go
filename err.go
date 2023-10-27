package client

import (
	"fmt"
)

type Err struct {
	ReqId      string `json:"reqId"`
	ErrorCode  int    `json:"errorCode"`
	ErrorMsg   string `json:"errorMsg"`
	HttpStatus int    `json:"httpStatus"`
}

func (e Err) Error() string {
	return fmt.Sprintf("client error; reqId: '%s' errorCode: '%d' errorMsg: '%s'", e.ReqId, e.ErrorCode, e.ErrorMsg)
}
