package sms

import (
	"time"

	"github.com/iamgoroot/kyivstar-opentelecom-go/internal/models"
)

type SendReq struct {
	From string `json:"from"`
	To   string `json:"to"`
	Text string `json:"text"`
}

type SendResp struct {
	models.ReqInfoGetter
	ReqID string `json:"reqId"`
	MsgID string `json:"msgId"`
}

type CheckResp struct {
	models.ReqInfoGetter
	ReqID  string     `json:"reqId"`
	MsgID  string     `json:"msgId"`
	Status string     `json:"status"`
	Date   *time.Time `json:"date"`
}

type BatchSendReq struct {
	Data map[string]SendReq `json:"data"`
}

type BatchSendItemResp struct {
	MsgID               string `json:"msgId,omitempty"`
	ReservedSmsSegments int    `json:"reservedSmsSegments,omitempty"`
	ErrorMsg            string `json:"errorMsg,omitempty"`
	ErrorCode           string `json:"errorCode,omitempty"`
}

type BatchSendResp struct {
	models.ReqInfoGetter
	ReqID string                       `json:"reqId"`
	Data  map[string]BatchSendItemResp `json:"data"`
}

type BatchStatusReq struct {
	Data []string `json:"data"`
}

type BatchStatusItemResp struct {
	MsgID  string `json:"msgId"`
	Status string `json:"status"`
	Date   string `json:"date"`
}

type BatchStatusResp struct {
	models.ReqInfoGetter
	ReqID string                         `json:"reqId"`
	Data  map[string]BatchStatusItemResp `json:"data"`
}
